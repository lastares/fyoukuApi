package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["fyoukuApi/controllers:BarrageController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:BarrageController"],
        beego.ControllerComments{
            Method: "BarrageSave",
            Router: "/barrage/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:BarrageController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:BarrageController"],
        beego.ControllerComments{
            Method: "BarrageWs",
            Router: "/barrage/ws",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:BarrageController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:BarrageController"],
        beego.ControllerComments{
            Method: "TestBarrages",
            Router: "/test/barrages",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:BaseController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:BaseController"],
        beego.ControllerComments{
            Method: "ChannelRegion",
            Router: "/channel/region",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:BaseController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:BaseController"],
        beego.ControllerComments{
            Method: "ChannelType",
            Router: "/channel/type",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:CommonController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:CommonController"],
        beego.ControllerComments{
            Method: "CommentList",
            Router: "/comment/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:CommonController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:CommonController"],
        beego.ControllerComments{
            Method: "CommentSave",
            Router: "/comment/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:TopController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:TopController"],
        beego.ControllerComments{
            Method: "ChannelTop",
            Router: "/channel/top",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:TopController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:TopController"],
        beego.ControllerComments{
            Method: "TypeTop",
            Router: "/type/top",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:UserController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserRegister",
            Router: "/register/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:UserController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/user/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelAdvert",
            Router: "/channel/advert",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelHotList",
            Router: "/channel/hot",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelRegionCommendVideo",
            Router: "/channel/recommend/region",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelTypeRecommendList",
            Router: "/channel/recommend/type",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelVideo",
            Router: "/channel/video",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "SendMessage",
            Router: "/send/message",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "UserVideo",
            Router: "/user/video",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoEpisodesList",
            Router: "/video/episodes/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoInfo",
            Router: "/video/info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoSave",
            Router: "/video/save",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuApi/controllers:WsController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:WsController"],
        beego.ControllerComments{
            Method: "TestWs",
            Router: "/test/ws",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
