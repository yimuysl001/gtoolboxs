package icons

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/constant"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/utils"
)

const (
	SAFARI_PINNED_TAB_SVG_RELATIVE_PATH = constant.ROOT_PATH + "/img/icons/safari-pinned-tab.svg"
	// 文件内容的16进制表示
	SAFARI_PINNED_TAB_SVG_HEX_CONTENT = `3c3f786d6c2076657273696f6e3d22312e3022207374616e64616c6f6e653d226e6f223f3e0d0a3c21444f435459504520737667205055424c494320222d2f2f5733432f2f445444205356472032303031303930342f2f454e220d0a2022687474703a2f2f7777772e77332e6f72672f54522f323030312f5245432d5356472d32303031303930342f4454442f73766731302e647464223e0d0a3c7376672076657273696f6e3d22312e302220786d6c6e733d22687474703a2f2f7777772e77332e6f72672f323030302f737667220d0a2077696474683d2231362e303030303030707422206865696768743d2231362e3030303030307074222076696577426f783d223020302031362e3030303030302031362e303030303030220d0a207072657365727665417370656374526174696f3d22784d6964594d6964206d656574223e0d0a3c6d657461646174613e0d0a4372656174656420627920706f747261636520312e31312c207772697474656e2062792050657465722053656c696e67657220323030312d323031330d0a3c2f6d657461646174613e0d0a3c67207472616e73666f726d3d227472616e736c61746528302e3030303030302c31362e30303030303029207363616c6528302e3030303332302c2d302e30303033323029220d0a66696c6c3d222330303030303022207374726f6b653d226e6f6e65223e0d0a3c7061746820643d224d313820343636313820633435202d373520313232202d32303720313232202d3231312030202d32203235202d3435203535202d3935203330202d35302035350d0a2d3936203535202d3130322030202d352035202d3130203130202d313020362030203130202d34203130202d392030202d35203733202d31333520313631202d323838203839202d3135330d0a313733202d32393820313837202d333233203134202d3235203332202d3537203431202d3732203838202d31343920313837202d33323420313839202d3333352032202d372038202d31330d0a3133202d3133203520302039202d342039202d31302030202d35203436202d383920313033202d31383720313735202d33303220343930202d38343620353037202d3837362038202d31360d0a3230202d3336203235202d3435203238202d343620323930202d34393820333339202d353835203133202d3233203734202d31323920313336202d323336203631202d313037203132330d0a2d32313520313337202d323430203134202d3235203239202d3530203333202d35362035202d35203233202d3337203430202d3730203138202d3333203338202d3637203434202d37350d0a3131202d3136203231202d3333203633202d313039203134202d3235203239202d3530203333202d35362034202d35203231202d3335203338202d3635203535202d313030203236310d0a2d34353520323639202d3436352034202d35203134202d3231203230202d3335203135202d3239203431202d373520313033202d313830203234202d3431203532202d38382036300d0a2d3130352039202d3136203537202d31303020313037202d31383520313132202d31393320333632202d36323620333830202d3636302038202d3134203233202d3338203333202d35350d0a3131202d3136203233202d3337203237202d34352034202d38203236202d3436203438202d3835203233202d3338203533202d3930203637202d313135203436202d38312036340d0a2d31313320313738202d333130203632202d31303720313231202d32313020313332202d323237203337202d3637203536202d3939203835202d313438203136202d3237203332202d35370d0a3336202d36352034202d38203135202d3237203235202d34322039202d3135203533202d3839203936202d313635203434202d373620313737202d33303720323936202d353133203132300d0a2d32303620323638202d34363320333330202d35373020313331202d32323720313137202d32303320323030202d333438203336202d3632203733202d313235203832202d3134302031300d0a2d3135203231202d3334203235202d34322034202d38203230202d3337203336202d3635203137202d3237203338202d3635203438202d3832203439202d3835203634202d3131312038370d0a2d313533203133202d3235203238202d3439203332202d35352034202d35203738202d31333420313635202d323835203837202d31353120313636202d32383820313736202d3330350d0a3130202d3136203236202d3433203335202d35392039202d313720313235202d32313720323537202d34343520313332202d32323920323533202d34343120323730202d3437312031370d0a2d3330203435202d3739203634202d313038203138202d3239203333202d3534203333202d35372030202d32203230202d3337203434202d3737203234202d343020313233202d3231320d0a323231202d333833203937202d31373020313930202d33333020323035202d333535203136202d3235203339202d3635203533202d3930203133202d3235203831202d313434203135320d0a2d323635203730202d31323120313337202d32333820313530202d323630203132202d3232203337202d3635203535202d3935203138202d3330203433202d3733203535202d39352031320d0a2d3232203438202d3835203830202d313430203737202d31333220313633202d32383020313930202d333330203133202d3232203731202d31323320313330202d3232352035390d0a2d31303220313136202d31393920313236202d323137203130202d3137203239202d3530203433202d3732203135202d3232203236202d3433203236202d34352030202d322032370d0a2d3530203630202d313036203333202d3536203630202d313033203630202d3130352030202d32203535202d3938203930202d3135352038202d313420313832202d333136203233390d0a2d343134203133202d3232203435202d3739203732202d313234203237202d3436203439202d3836203439202d38392030202d32203134202d3234203330202d3438203136202d32340d0a3330202d3436203330202d34392030202d35203734202d31333520313030202d3137362035202d38203234202d3432203433202d3735203530202d3838203538202d313031203236320d0a2d34353520313034202d31373920313939202d33343520323133202d333730203134202d3235203238202d3439203332202d35352034202d35203137202d3236203238202d34352031300d0a2d3139203632202d31303920313134202d32303020313134202d31393720313333202d32333020313730202d323935203136202d3237203333202d3537203338202d3635203137202d32380d0a3936202d31363520313033202d3138302034202d38203136202d3238203236202d3435203130202d3136203737202d31333120313438202d323535203732202d313234203138310d0a2d33313320323433202d343230203632202d31303720313231202d32303920313331202d323237203335202d363220333233202d35363020333932202d363738203338202d36362038330d0a2d31343520313030202d313735203136202d3330203333202d3539203337202d36352034202d35203137202d3237203239202d3437203334202d3631203536202d3130302039300d0a2d313536203137202d3239203331202d3535203331202d35372030202d32203137202d3332203339202d3637203231202d333520313334202d32323920323531202d343333203131370d0a2d32303320323335202d34303720323631202d343531203237202d3435203439202d3835203439202d38382030202d342038202d3139203139202d3334203135202d3231203230300d0a2d33343120333039202d353333203130202d3139203333202d3538203531202d3837203137202d3239203331202d3534203331202d35362030202d32203235202d3434203535202d39340d0a3330202d3530203535202d3935203535202d39382030202d342036202d3135203134202d32332037202d39203237202d3431203433202d3731203137202d333020313730202d3239370d0a333432202d35393420313731202d32393620333131202d35343220333131202d3534372030202d352035202d39203130202d3920362030203130202d34203130202d31302030202d350d0a3232202d3437203439202d3932203237202d3436203538202d3939203638202d313138203234202d3433203831202d313430203933202d3136302035202d38203636202d313134203133350d0a2d323335203639202d31323120313330202d32323720313335202d323335203132202d323120323539202d34343720323833202d343930203130202d3139203238202d34372033380d0a2d3632203131202d3134203139202d3239203139202d33322030202d33203337202d3639203833202d313438203939202d31373020333035202d35323620333337202d3538332031330d0a2d3232203331202d3533203431202d3730203131202d3136203232202d3337203236202d34352037202d3134203832202d31343620313033202d313830203134202d3234203138310d0a2d33313120323035202d333535203133202d3232203436202d3830203735202d313330203239202d3439203634202d313130203738202d313335203134202d3235203531202d38382038320d0a2d313430203331202d3532203539202d313032203633202d3131302034202d38203138202d3333203331202d353520323035202d33353320323834202d34383920333039202d3533350d0a3137202d3330203435202d3738203632202d313036203138202d3238203336202d3630203339202d37322034202d3132203132202d3232203137202d3232203520302039202d3420390d0a2d31302030202d3520313039202d31393720323431202d34323720313333202d32333020323530202d34333120323539202d343438203531202d393020323232202d333835203238300d0a2d343835203337202d3633203738202d313335203932202d313630203134202d3235203637202d31313720313138202d323035203531202d383820313031202d313735203131310d0a2d313933203334202d3538203535202d393520313439202d323537203531202d383820313031202d31373320313130202d3139302039202d3136203736202d31333120313437202d3235350d0a3732202d31323420313430202d32343120313531202d323630203631202d31303820323831202d34383920333535202d363135203338202d3636203737202d313333203837202d3135300d0a3335202d3633203931202d31363120313030202d313735203134202d3233203939202d31363920313238202d323230203534202d393720313335202d32333520313432202d32343520340d0a2d35203230202d3332203335202d3630203236202d343820323338202d34313620323736202d343830203130202d3136203236202d3436203337202d3635203330202d3533203338320d0a2d36363120343033202d363935203130202d3136203232202d3337203236202d34352034202d38203236202d3438203530202d3838203234202d3431203433202d3735203433202d37370d0a30202d32203232202d3430203530202d3835203237202d3435203530202d3834203530202d38362030202d33203338202d3639203833202d313437203834202d313432203330320d0a2d35323020333430202d353837203130202d3139203334202d3630203532202d3930203138202d3330203434202d3735203537202d313030203134202d3235203435202d37392037300d0a2d313230203235202d3431203536202d3936203730202d313231203134202d3235203737202d31333320313338202d323430203632202d31303720313232202d323130203133320d0a2d323239203235202d343320333130202d35333520333337202d353831203131202d3139203236202d3435203334202d3539203137202d333220323338202d34313420323636202d3436300d0a3131202d3139203234202d3431203238202d34392033202d37203735202d31333320313630202d323738203834202d31343620313533202d32363920313533202d3237342030202d3520350d0a2d39203130202d3920362030203130202d34203130202d31302030202d35203832202d31353020313831202d33323220313832202d33313420323031202d33343620323430202d3431350d0a3132202d3231203830202d31333920313532202d323633203731202d31323420313431202d32343520313535202d323730203134202d3235203238202d3439203332202d35352036202d380d0a313435202d32343820323230202d333830203337202d363620323039202d33363220323239202d333935203131202d3139203234202d3432203238202d34392034202d382036370d0a2d31313820313430202d323433203733202d31323520313333202d32333020313333202d3233332030202d32203135202d3238203333202d3537203139202d3239203437202d37382036340d0a2d313038203137202d3330203533202d3933203739202d313339203533202d3930203832202d31343120313537202d323732203832202d31343220313135202d313939203338310d0a2d36353920313432202d32343520323638202d34363320323831202d343835203132202d3232203731202d31323520313332202d323330203630202d31303420313732202d3239380d0a323438202d343330203736202d31333220313436202d32353320313536202d323730203131202d3136203232202d3336203236202d34342033202d38203330202d3534203630202d3130330d0a3239202d3439203533202d3931203533202d39332030202d33203138202d3334203430202d3730203232202d3336203430202d3637203430202d36392030202d32203337202d36362038310d0a2d313432203435202d3737203938202d31363820313139202d323034203230202d3336203437202d3831203538202d313030203132202d3139203237202d3437203333202d363220360d0a2d3136203135202d3238203230202d3238203520302039202d342039202d392030202d36203633202d31313820313430202d323531203737202d31333320313430202d323433203134300d0a2d3234352030202d32203138202d3333203431202d3730203232202d3337203439202d3833203630202d313031203130202d3139203239202d3531203430202d3731203235202d34350d0a313039202d31383920313236202d3231382037202d3131203137202d3239203232202d34302036202d3131203232202d3338203335202d3630203134202d3232203337202d36322035320d0a2d3930203134202d3237203335202d3632203435202d3737203131202d3134203139202d3239203139202d33322030202d33203138202d3335203430202d3731203232202d33362034300d0a2d3637203430202d36392030202d32203139202d3335203432202d3732203233202d3338203535202d3934203732202d313234203236202d343720313339202d323434203137310d0a2d3239382036202d39203231202d3336203334202d3630203238202d3438203337202d3531203531202d313920362031322031392033362032392035322031302031372032372034360d0a333820363520313120313920313034203138312032303820333630203130332031373920313939203334352032313320333730203134203235203432203734203634203130392032310d0a33342033382036352033382036372030203220313820333320343020363920323220333620343020363720343020363920302033203137372033313020313939203334362031362032360d0a31333620323334203134302032343420322035203235203434203532203838203237203434203439203831203439203834203020322031382033342034302037302032322033362034300d0a363720343020363920302032203230203336203433203737203335203538203136392032383920323937203531332039203137203530203836203930203135352034302036392038360d0a313530203130332031383020313620333020333520363220343120373020362038203136203234203232203335203335203634203732203132392031363720323933203539203130300d0a3131362031393920313237203232302031312032302033302035332034312037322034332037322031303730203138353020313132312031393430203134203235203635203131330d0a31313320313935203438203833203936203136362031303720313835203130203139203238203530203338203638203131203138203733203132342031333720323335203634203131310d0a3137352033303320323436203432372037312031323420313733203239392032323520333930203532203931203131362032303220313433203234382032372034352034392038352034390d0a383920302034203620313420313420323220372039203238203433203436203736203236203437203235312034333620333738203635352031312031392032392035312034302037300d0a31312031392031303120313736203230312033343820393920313732203138312033313720313831203332332030203520352039203130203920362030203130203520313020313120300d0a3620382032332031382033372031312031352033322035322034392038322031362033302031333020323238203235332034343020313232203231322032333420343035203234380d0a34333020313320323520333920373020353720313030203339203635203639203131372031333020323235203235203434203530203837203535203935203132203139203738203133340d0a323230203338302036312031303720313239203232342031353020323630203136312032373720323232203338322032343620343235203135203238203437203833203731203132330d0a323420343120343320373820343320383320302035203420392038203920342030203133203132203139203238203720313520323320343520333620363720363620313130203237370d0a3437382032373720343833203020332036203133203134203231203720392032372034312034332037312031372033302034352038302036332031313020333420353720333735203634390d0a333934203638352036203131203136203237203232203335203620382032362034322034342037352031382033332034312037342035312039302031302031372032342034312033320d0a35352035342039372037322031323820383820313532203131203134203139203238203139203330203020332037392031343120313735203330382039362031363720313735203330350d0a3137352033303820302033203620313320313420323120372039203236203339203431203636203333203630203237362034383320333338203538372032342034302034362038302035300d0a38382034203820313320323420323020333520313420323320393520313633203132352032313520313120313920353220393120393220313630203430203639203830203133392039300d0a31353520392031372031303320313739203230372033363020313035203138322032303020333436203231312033363520313033203138312034363320383032203438392038343520370d0a31312031352032372031392033352034203820323920353120353520393520363420313130203832382031343333203834382031343730203920313720323420343120333320353520390d0a3134203239203438203435203737203135203238203532203933203832203134352033302035312036322031303720373120313233203137203330203233312033393820343030203639300d0a3531203838203130332031373920313135203230322031322032332032362034382033322035352036203720323420333820343020363820313720333020363120313037203938203137300d0a33372036332038342031343420313033203138302031392033362034312037322034382038312038203820313420313820313420323120302034203237203531203539203130362033320d0a35352037322031323420383920313534203136203239203731203132352031323220323133203531203838203130342031383020313138203230352031332032352032382035302033320d0a35352034203620313720323620323820343520313120313920343520383020373720313335203331203535203636203131362037372031333520313120313920383820313532203137310d0a3239352034303120363934203632302031303732203635302031313235203131203139203837203135322031373020323935203833203134332031353820323733203136362032383820390d0a313620323120333620323620343520362039203331203532203535203936203235203433203534203934203636203131352031312032302039352031363420313836203332312039310d0a3135372031373320323939203138322033313520392031372032362034362033372036352031322031392036362031313420313231203231302035362039362031303820313836203131370d0a3230302038203134203234203430203334203539203234203435203338332036363420343132203731332035203920313720323920323620343520313520323820313230203231300d0a32343120343139203336203631203638203131372037322031323520342038203132203233203139203334203335203537203234352034323020323632203435332031312032302033350d0a3631203533203930203137203239203332203534203332203536203020332032382035312036322031303820333320353720373020313139203830203133382031302031392032332034320d0a32382035302035203820333220353320353920313030203237203437203134392032353820323731203437302031323220323132203233342034303520323438203433302033302035330d0a3632203130382038302031333520362031312031352032372031392033352034203820383520313530203138312033313520393620313635203138372033323320323032203335302033310d0a3536203131362032303220313330203232352035203820323520343220343320373520313920333320393220313539203136322032383020313439203235372031353720323731203230320d0a33353020313920333320333820363720343320373520392031342032323820333932203237352034373520313220323220353520393620393520313635203430203639203830203133390d0a393020313535203234203432203230322033353020323231203338332039203135203237203437203431203732203134203235203735203133312031333620323336203631203130360d0a3132312032313020313334203233322039392031373220323731203437302032373920343832203520382032332034302034302037302031382033302038312031343120313432203234350d0a36302031303520313231203231302031333520323335203134203235203731203132342031323720323230203536203936203134332032343720313934203333352035312038382039360d0a313637203130322031373520313420323420313830203331312032303420333535203233203433203334302035393020333536203631352035203820353020383720313031203137350d0a313731203330312035313720383938203538322031303038203235203433203436203831203436203833203020322031322032332032372034372031342032332034302036372035360d0a3937203136203330203335203632203432203730203720382031352032322031382033302034203820323020333820333720363520313620323820333320353720333720363520362031320d0a3131312031393620313433203235302035203820353520393520313132203139332035372039382031313320313935203132362032313520313220323020323720343620333220353720360d0a3131203134203237203230203335203520382037362031333020313536203237302038302031343020313635203238372031383720333235203233203339203532203930203636203131350d0a3133203235203330203532203337203631203820382031342031382031342032312030203420343120373720393220313635203530203837203137352033303220323736203437380d0a31303120313736203230382033363020323336203430382032382034392036372031313720383620313532203139203335203431203730203438203737203620362031322031352031320d0a3139203020372031323420323234203136372032393120313220323120323320343020323320343220302032203231203430203436203833203236203433203535203932203634203130390d0a35342039352033323720353638203335342036313420313920333020343520373520353920313030203731203132382038322031343520383920313438203420322038203820382031330d0a30203520343220383220393420313732203331312035333820343936203835382035313820383937203134203235203430203730203538203130302031382033302034322037312035330d0a39302031302031392037392031333920313532203236352037332031323720313432203234362031353320323635203130203139203433203736203732203132352032392035302036330d0a31303820373520313330203635203131362038302031343020383720313433203420322038203820382031322030203820313134203231322031343020323530203620382031342032340d0a323020333520352031312035342039372031303820313930206c31303020313730202d39363131203320632d353238362031202d39363134202d31202d39363138202d35202d35202d360d0a2d343139202d373139202d363139202d31303638202d3839202d313535202d323637202d343633202d333233202d353630202d3338202d3636202d3831202d313430202d3935202d3136350d0a2d3331202d3536202d323633202d343537202d353236202d393130202d313130202d313930202d323234202d333838202d323534202d343430202d3239202d3532202d3631202d3130390d0a2d3731202d313235202d3233202d3339202d323433202d343230202d323638202d343635202d3131202d3139202d323034202d333532202d343238202d373430202d323234202d3338380d0a2d343737202d383236202d353633202d393735202d3835202d313438202d313835202d333232202d323232202d333835202d3337202d3633202d313230202d323037202d3138350d0a2d333230202d3635202d313133202d313737202d333036202d323438202d343330202d3732202d313234202d313732202d323937202d323232202d333835202d3531202d3838202d3134320d0a2d323435202d323032202d333530202d313331202d323236202d323437202d343237202d343038202d373035202d3635202d313133202d323439202d343332202d343130202d3731300d0a2d313630202d323738202d333838202d363733202d353036202d383737202d313138202d323035202d323136202d333733202d323139202d333733202d332030202d35322038320d0a2d31303920313833202d353820313030202d31343420323530202d31393220333332202d393520313634202d34303220363936202d3634372031313230202d383520313439202d3232380d0a333936202d33313720353530202d32313220333635202d3938322031373030202d313030382031373435202d3130203139202d3433203736202d373220313235202d3239203530202d36340d0a313130202d373720313335202d3134203235202d363320313130202d31313020313930202d3437203830202d393620313635202d31313020313930202d3134203235202d3939203137310d0a2d31383820333235202d383920313534202d31373420333030202d31383820333235202d3133203235202d363420313133202d31313220313935202d3438203833202d313430203234320d0a2d32303520333535202d363520313133202d31383320333137202d32363320343534202d373920313337202d31353220323634202d31363320323832202d3530203839202d3333350d0a353833202d33353420363134202d3132203139202d3334203538202d3530203835202d3135203238202d31323920323236202d32353320343430202d31323420323135202d3233350d0a343038202d32343720343330202d3132203232202d363920313231202d31323720323230202d3538203939202d32323620333839202d33373320363435202d31343820323536202d3332340d0a353631202d33393220363738202d363720313137202d31333420323332202d31343720323535202d3133203233202d3333203539202d3436203830206c2d3232203337202d3936313520300d0a2d393631352030203230202d33327a222f3e0d0a3c2f673e0d0a3c2f7376673e0d0a`
)

func AddRouterOfSafariPinnedTabSvg(router *ghttp.RouterGroup) {

	utils.GetOther(router, SAFARI_PINNED_TAB_SVG_RELATIVE_PATH, SAFARI_PINNED_TAB_SVG_HEX_CONTENT)

}