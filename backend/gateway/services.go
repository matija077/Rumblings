package main

import "net/http"

type port = uint16
type serviceName = string
type serviceInternalName = string
type route = string

type registry map[serviceName]port

type service struct {
	port
	serviceInternalName
	route
}

type Backend struct {
	service
}
type Backends map[serviceName]Backend

type Frontend struct {
	service
	backends []serviceName
}
type Frontends map[serviceName]Frontend

func (f *Frontend) getBackendRoutes() []route {
	return f.backends
}

type Services map[route]serviceName

type createServicesModelReturnType struct {
	getFrontend func(route route) (*Frontend, bool)
	redirect    func(w *http.ResponseWriter, r *http.Request, f *Frontend)
}

func createServicesModel() createServicesModelReturnType {
	var backends Backends = make(Backends, 2)
	var frontends Frontends = make(Frontends, 2)

	backends["backend1"] = Backend{
		service: service{
			port:                9010,
			serviceInternalName: "backend1",
		},
	}

	backends["backend2"] = Backend{
		service: service{
			port:                9020,
			serviceInternalName: "backend2",
		},
	}

	frontends["main"] = Frontend{
		service: service{
			port:                2010,
			serviceInternalName: "frontend",
		},
		backends: []serviceName{"backend1", "backend2"},
	}

	frontends["showLogs"] = Frontend{
		service: service{
			port:                2020,
			serviceInternalName: "logger",
		},
		backends: []serviceName{},
	}

	var services Services = make(Services, 2)
	services["/"] = "main"
	services["/log"] = "showLogs"

	return createServicesModelReturnType{
		getFrontend: func(route route) (*Frontend, bool) {
			serviceName, ok := services[route]
			if !ok {
				return nil, false
			}
			if frontend, ok := frontends[serviceName]; ok {
				return &frontend, true
			}

			return nil, false
		},
		redirect: func(w *http.ResponseWriter, r *http.Request, f *Frontend) {
			http.Redirect(*w, r, "127.0.0.1/"+f.route, http.StatusSeeOther)
		},
	}
}
