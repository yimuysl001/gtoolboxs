package swagger

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/knife"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/knife/img/icons"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/knife/oauth"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/knife/webjars/css"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/knife/webjars/fonts"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/knife/webjars/img"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/knife/webjars/js"
)

// InitSwaggerByStatic
//
//	@Description:
//	@param router
func InitSwaggerByStatic(router *ghttp.RouterGroup) {
	knife.AddApiDocRouter(router)
	knife.AddSwaggerResourcesRouter(router)

}

// InitSwaggerKnife
//
//	@Description:
//	@param router
func InitSwaggerKnife(router *ghttp.RouterGroup) {

	knife.AddApiDocRouter(router)
	knife.AddSwaggerResourcesRouter(router)

	knife.AddRouterOfDocHtml(router)

	knife.AddRouterOfFaviconIco(router)

	icons.AddRouterOfAndroidChrome192x192Png(router)

	icons.AddRouterOfAndroidChrome512x512Png(router)

	icons.AddRouterOfAppleTouchIcon120x120Png(router)

	icons.AddRouterOfAppleTouchIcon152x152Png(router)

	icons.AddRouterOfAppleTouchIcon180x180Png(router)

	icons.AddRouterOfAppleTouchIcon60x60Png(router)

	icons.AddRouterOfAppleTouchIcon76x76Png(router)

	icons.AddRouterOfAppleTouchIconPng(router)

	icons.AddRouterOfFavicon16x16Png(router)

	icons.AddRouterOfFavicon32x32Png(router)

	icons.AddRouterOfMsapplicationIcon144x144Png(router)

	icons.AddRouterOfMstile150x150Png(router)

	icons.AddRouterOfSafariPinnedTabSvg(router)

	oauth.AddRouterOfJqueryMinJs(router)

	oauth.AddRouterOfJqueryMinJsGz(router)

	oauth.AddRouterOfOauth2Html(router)

	knife.AddRouterOfRobotsTxt(router)

	css.AddRouterOfAppB848c085Css(router)

	css.AddRouterOfAppB848c085CssGz(router)

	css.AddRouterOfChunk51277dbe57225f85Css(router)

	css.AddRouterOfChunkVendors3f2387deCss(router)

	css.AddRouterOfChunkVendors3f2387deCssGz(router)

	fonts.AddRouterOfFontawesomeWebfont706450d7Ttf(router)

	fonts.AddRouterOfFontawesomeWebfont97493d3fWoff2(router)

	fonts.AddRouterOfFontawesomeWebfontD9ee23d5Woff(router)

	fonts.AddRouterOfFontawesomeWebfontF7c2b4b7Eot(router)

	fonts.AddRouterOfIconfont4ca3d0c0Ttf(router)

	fonts.AddRouterOfIconfontE2d2b98eEot(router)

	img.AddRouterOfEditormdLogo84b6c2a9Svg(router)

	img.AddRouterOfFontawesomeWebfont139e74e2Svg(router)

	img.AddRouterOfIconfontDd63dc33Svg(router)

	img.AddRouterOfLoadingC929501eGif(router)

	img.AddRouterOfLoading2x695405a9Gif(router)

	img.AddRouterOfLoading3x65eacf61Gif(router)

	js.AddRouterOfApp2650dddfJs(router)

	js.AddRouterOfApp2650dddfJsGz(router)

	js.AddRouterOfChunk069eb437371ae4fdJs(router)

	js.AddRouterOfChunk069eb437371ae4fdJsLICENSETxt(router)

	js.AddRouterOfChunk069eb437371ae4fdJsGz(router)

	js.AddRouterOfChunk0fd67716D57e2c41Js(router)

	js.AddRouterOfChunk0fd67716D57e2c41JsGz(router)

	js.AddRouterOfChunk2d0af44eC299c1d4Js(router)

	js.AddRouterOfChunk2d0bd799Cc91c520Js(router)

	js.AddRouterOfChunk2d0d0b98Cb1dea78Js(router)

	js.AddRouterOfChunk2d0da532Dd3c929cJs(router)

	js.AddRouterOfChunk2d22269dBd9173e1Js(router)

	js.AddRouterOfChunk3b888a658737ce4fJs(router)

	js.AddRouterOfChunk3b888a658737ce4fJsGz(router)

	js.AddRouterOfChunk3ec4aaa8A79d19f8Js(router)

	js.AddRouterOfChunk3ec4aaa8A79d19f8JsGz(router)

	js.AddRouterOfChunk51277dbeA577fa2fJs(router)

	js.AddRouterOfChunk51277dbeA577fa2fJsLICENSETxt(router)

	js.AddRouterOfChunk51277dbeA577fa2fJsGz(router)

	js.AddRouterOfChunk589faee0B24e5f3dJs(router)

	js.AddRouterOfChunk589faee0B24e5f3dJsLICENSETxt(router)

	js.AddRouterOfChunk589faee0B24e5f3dJsGz(router)

	js.AddRouterOfChunk735c675c76ef1019Js(router)

	js.AddRouterOfChunk735c675c76ef1019JsGz(router)

	js.AddRouterOfChunkAdb9e944B888f4bdJs(router)

	js.AddRouterOfChunkAdb9e944B888f4bdJsLICENSETxt(router)

	js.AddRouterOfChunkAdb9e944B888f4bdJsGz(router)

	js.AddRouterOfChunkVendors90e8ba20Js(router)

	js.AddRouterOfChunkVendors90e8ba20JsLICENSETxt(router)

	js.AddRouterOfChunkVendors90e8ba20JsGz(router)

}
