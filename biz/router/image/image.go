package image

import (
	image "qnc/biz/handler/image"

	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_image := root.Group("/image", _imageMw()...)
		{
			_ud := _image.Group("/ud", _udMw()...)
			_ud.POST("/", append(_imageudMw(), image.ImageUd)...)
		}
		{
			_ud := _image.Group("/tryon", _udMw()...)
			_ud.POST("/", append(_imageudMw(), image.ImageTryOn)...)
		}
	}
}
