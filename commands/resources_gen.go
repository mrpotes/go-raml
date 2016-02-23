package commands

import (
	"strings"

	"github.com/Jumpscale/go-raml/raml"
	log "github.com/Sirupsen/logrus"
)

const (
	maxCommentPerLine = 80
)

// resourceDef is Go code representation of a resource
type resourceDef struct {
	Name        string            // resource name
	Endpoint    string            // root endpoint
	Methods     []interfaceMethod // all methods of this resource
	IsServer    bool              // true if it is resource definition for server
	NeedJSON    bool              // if true, the API implementation to import encoding/json package
	PackageName string            // Name of the package this resource resides in
}

// create a resource definition
func newResourceDef(endpoint, packageName string) resourceDef {
	rd := resourceDef{
		Endpoint: endpoint,
	}
	rd.Name = strings.Title(normalizeURI(endpoint))
	rd.PackageName = packageName
	return rd
}

// method of resource's interface
type interfaceMethod struct {
	*raml.Method
	MethodName   string
	Endpoint     string
	Verb         string
	ReqBody      string         // request body type
	RespBody     string         // response body type
	ResourcePath string         // normalized resource path
	Resource     *raml.Resource // resource object of this method
	Params       string         // methods params
	FuncComments []string
}

// create an interfaceMethod object
func newInterfaceMethod(r *raml.Resource, rd *resourceDef, m *raml.Method, methodName, parentEndpoint, curEndpoint string) interfaceMethod {
	im := interfaceMethod{
		Method:   m,
		Endpoint: parentEndpoint + curEndpoint,
		Verb:     strings.ToUpper(methodName),
		Resource: r,
	}

	// set request body
	im.ReqBody = assignBodyName(m.Bodies, normalizeURITitle(im.Endpoint)+methodName, "ReqBody")

	//set response body
	for k, v := range m.Responses {
		if k >= 200 && k < 300 {
			im.RespBody = assignBodyName(v.Bodies, normalizeURITitle(im.Endpoint)+methodName, "RespBody")
		}
	}

	// if there is request/response body, then it needs to import encoding/json
	if im.RespBody != "" || im.ReqBody != "" {
		rd.NeedJSON = true
	}

	// set func comment
	if len(m.Description) > 0 {
		im.FuncComments = funcCommentBuilder(m.Description)
	}

	return im
}

func newServerInterfaceMethod(r *raml.Resource, rd *resourceDef, m *raml.Method, methodName, parentEndpoint, curEndpoint, lang string) interfaceMethod {
	im := newInterfaceMethod(r, rd, m, methodName, parentEndpoint, curEndpoint)

	name := normalizeURI(im.Endpoint)
	if lang == "go" {
		im.MethodName = name[len(rd.Name):] + methodName
	} else {
		im.MethodName = snakeCaseResourceURI(r) + "_" + strings.ToLower(im.Verb)
		im.Params = strings.Join(getResourceParams(r), ", ")
		im.Endpoint = strings.Replace(im.Endpoint, "{", "<", -1)
		im.Endpoint = strings.Replace(im.Endpoint, "}", ">", -1)
	}

	return im
}

func newClientInterfaceMethod(r *raml.Resource, rd *resourceDef, m *raml.Method, methodName, parentEndpoint, curEndpoint string) (interfaceMethod, error) {
	im := newInterfaceMethod(r, rd, m, methodName, parentEndpoint, curEndpoint)

	// build func/method params
	postBuildParams := func(r *raml.Resource, bodyType string) (string, error) {
		paramsStr := strings.Join(getResourceParams(r), ",")
		if len(paramsStr) > 0 {
			paramsStr += " string"
		}

		// append request body type
		if len(bodyType) > 0 {
			if len(paramsStr) > 0 {
				paramsStr += ", "
			}
			paramsStr += strings.ToLower(bodyType) + " " + bodyType
		}

		// append header
		if len(paramsStr) > 0 {
			paramsStr += ","
		}
		paramsStr += "headers,queryParams map[string]interface{}"

		return paramsStr, nil
	}

	im.ResourcePath = paramizingURI(im.Endpoint)

	name := normalizeURITitle(im.Endpoint)
	im.MethodName = strings.Title(name + methodName)

	im.ReqBody = assignBodyName(m.Bodies, name+methodName, "ReqBody")

	methodParam, err := postBuildParams(r, im.ReqBody)
	if err != nil {
		return im, err
	}
	im.Params = methodParam

	return im, nil
}

// assignBodyName assign bodies by bodies.Type or bodies.ApplicationJson
// if bodiesType generated from bodies.Type we dont need append prefix and suffix
// 		example : bodies.Type = City, so bodiesType = City
// if bodiesType generated from bodies.ApplicationJson, we get that value from prefix and suffix
//		suffix = [ReqBody | RespBody] and prefix should be uri + method name.
//		example prefix could be UsersUserIdDelete
func assignBodyName(bodies raml.Bodies, prefix, suffix string) string {
	var bodiesType string

	if len(bodies.Type) > 0 {
		bodiesType = bodies.Type
	} else if bodies.ApplicationJson != nil {
		if bodies.ApplicationJson.Type != "" {
			bodiesType = bodies.ApplicationJson.Type
		} else {
			bodiesType = prefix + suffix
		}
	}

	return bodiesType
}

// add a method to resource definition
func (rd *resourceDef) addMethod(r *raml.Resource, m *raml.Method, methodName, parentEndpoint, curEndpoint, lang string) {
	if m == nil {
		return
	}
	var im interfaceMethod
	var err error
	if rd.IsServer {
		im = newServerInterfaceMethod(r, rd, m, methodName, parentEndpoint, curEndpoint, lang)
	} else {
		im, err = newClientInterfaceMethod(r, rd, m, methodName, parentEndpoint, curEndpoint)
		if err != nil {
			log.Errorf("client interface method error, err = %v", err)
			return
		}
	}
	rd.Methods = append(rd.Methods, im)
}

// generate all methods of a resource recursively
func (rd *resourceDef) generateMethods(r *raml.Resource, parentEndpoint, curEndpoint, lang string) {
	rd.addMethod(r, r.Get, "Get", parentEndpoint, curEndpoint, lang)
	rd.addMethod(r, r.Post, "Post", parentEndpoint, curEndpoint, lang)
	rd.addMethod(r, r.Put, "Put", parentEndpoint, curEndpoint, lang)
	rd.addMethod(r, r.Patch, "Patch", parentEndpoint, curEndpoint, lang)
	rd.addMethod(r, r.Delete, "Delete", parentEndpoint, curEndpoint, lang)

	for k, v := range r.Nested {
		rd.generateMethods(v, parentEndpoint+curEndpoint, k, lang)
	}
}