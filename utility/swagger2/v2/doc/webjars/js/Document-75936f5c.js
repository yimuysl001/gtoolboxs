import{u as utils,_ as _export_sfc,V as VAceEditor,d as defineAsyncComponent,a as __vitePreload,b as useGlobalsStore,c as computed,e as useknife4jModels,f as useI18n,g as constants,m as message,h as uniqueId,i as cloneDeep,r as resolveComponent,o as openBlock,j as createElementBlock,k as createBaseVNode,l as createVNode,w as withCtx,t as toDisplayString,n as normalizeClass,p as createTextVNode,q as createCommentVNode,F as Fragment,s as createBlock,v as renderList,x as __unplugin_components_5,y as __unplugin_components_1,z as __unplugin_components_5$1,A as __unplugin_components_3,T as Tabs}from"./doc-7814a93f.js";import{C as ClipboardJS}from"./clipboard-814f04b9.js";function markdownSingleText(t){var e=[];return t!=null&&t!=null&&(markdownLines$1(e),e.push("## "+t.summary),markdownLines$1(e),e.push("**接口地址**:`"+t.showUrl+"`"),markdownLines$1(e),e.push("**请求方式**:`"+t.methodType+"`"),markdownLines$1(e),e.push("**请求数据类型**:`"+utils.toString(t.consumes,"*")+"`"),markdownLines$1(e),e.push("**响应数据类型**:`"+utils.toString(t.produces,"*")+"`"),markdownLines$1(e),e.push("**接口描述**:"+utils.toString(t.description,"暂无")),utils.checkUndefined(t.requestValue)&&(markdownLines$1(e),e.push("**请求示例**:"),markdownLines$1(e),e.push("```javascript"),e.push(t.requestValue),e.push("```")),createApiRequestParameters$1(t,e),createApiResponseStatus$1(t,e),createApiResponseParameters$1(t,e)),e.join(`
`)}function markdownLines$1(t){t.push(`
`)}function createApiRequestParameters$1(t,e){let s=t.reqParameters;markdownLines$1(e),e.push("**请求参数**:"),markdownLines$1(e),e.push("**请求参数**:"),s.length>0?(markdownLines$1(e),e.push("| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |"),e.push("| -------- | -------- | ----- | -------- | -------- | ------ |"),deepMdTableByRequestParameter$1(s,e,1)):(markdownLines$1(e),e.push("暂无"))}function createApiResponseStatus$1(t,e){utils.checkUndefined(t.responseCodes)&&t.responseCodes.length>0&&(markdownLines$1(e),e.push("**响应状态**:"),markdownLines$1(e),e.push("| 状态码 | 说明 | schema |"),e.push("| -------- | -------- | ----- | "),t.responseCodes.forEach(function(s){e.push("|"+utils.toString(s.code,"")+"|"+utils.toString(s.description,"")+"|"+utils.toString(s.schema,"")+"|")}))}function createApiResponseParameters$1(t,e){if(t.multipartResponseSchema){var s=t.multipCodeDatas;utils.arrNotEmpty(s)&&s.forEach(function(n){markdownLines$1(e),e.push("**响应状态码-"+utils.toString(n.code,"")+"**:"),createApiResponseSingleParam$1(n,e)})}else createApiResponseSingleParam$1(t.multipData,e)}function createApiResponseSingleParam$1(t,e){createApiResponseHeaderParams$1(t.responseHeaderParameters,e),markdownLines$1(e),e.push("**响应参数**:"),markdownLines$1(e),utils.arrNotEmpty(t.data)?(e.push("| 参数名称 | 参数说明 | 类型 | schema |"),e.push("| -------- | -------- | ----- |----- | "),t.data.forEach(function(s){s.level=1,e.push("|"+getMdTableByLevel$1(s)+"|"+utils.toString(s.description,"")+"|"+utils.toString(s.type,"")+"|"+utils.toString(s.schemaValue,"")+"|"),deepMdTableByResponseParameter$1(s.children,e,s.level+1)})):e.push("暂无"),markdownLines$1(e),e.push("**响应示例**:"),t.responseBasicType?(e.push("```text"),e.push(t.responseText),e.push("```")):(e.push("```javascript"),e.push(t.responseValue),e.push("```"))}function deepMdTableByResponseParameter$1(t,e,s){t!=null&&t!=null&&t.length>0&&t.forEach(function(n){n.level=s,e.push("|"+getMdTableByLevel$1(n)+"|"+utils.toString(n.description,"")+"|"+utils.toString(n.type,"")+"|"+utils.toString(n.schemaValue,"")+"|"),deepMdTableByResponseParameter$1(n.children,e,n.level+1)})}function createApiResponseHeaderParams$1(t,e){utils.checkUndefined(t)&&t.length>0&&(markdownLines$1(e),e.push("**响应Header**:"),markdownLines$1(e),e.push("| 参数名称 | 参数说明 | 数据类型 |"),e.push("| -------- | -------- | ----- | "),t.forEach(function(s){e.push("|"+utils.toString(s.name,"")+"|"+utils.toString(s.description,"")+"|"+utils.toString(s.type,"")+"|")}))}function deepMdTableByRequestParameter$1(t,e,s){t!=null&&t!=null&&t.length>0&&t.forEach(function(n){n.level=s,e.push("|"+getMdTableByLevel$1(n)+"|"+utils.toString(n.description,"")+"|"+utils.toString(n.in,"")+"|"+utils.toString(n.require,"")+"|"+utils.toString(n.type,"")+"|"+utils.toString(n.schemaValue,"")+"|"),deepMdTableByRequestParameter$1(n.children,e,n.level+1)})}function getMdTableByLevel$1(t){for(var e=[],s=1;s<t.level;s++)e.push("&emsp;&emsp;");var n=e.join("")+t.name;return n}function markdownSingleTextUs(t){var e=[];return t!=null&&t!=null&&(markdownLines(e),e.push("## "+t.summary),markdownLines(e),e.push("**Url**:`"+t.showUrl+"`"),markdownLines(e),e.push("**Method**:`"+t.methodType+"`"),markdownLines(e),e.push("**produces**:`"+utils.toString(t.consumes,"*")+"`"),markdownLines(e),e.push("**consumes**:`"+utils.toString(t.produces,"*")+"`"),markdownLines(e),e.push("**description**:"+utils.toString(t.description,"None")),utils.checkUndefined(t.requestValue)&&(markdownLines(e),e.push("**Sample**:"),markdownLines(e),e.push("```javascript"),e.push(t.requestValue),e.push("```")),createApiRequestParameters(t,e),createApiResponseStatus(t,e),createApiResponseParameters(t,e)),e.join(`
`)}function markdownLines(t){t.push(`
`)}function createApiRequestParameters(t,e){let s=t.reqParameters;markdownLines(e),e.push("**Params**:"),markdownLines(e),e.push("**Params**:"),s.length>0?(markdownLines(e),e.push("| name | description | in    | require | type | schema |"),e.push("| -------- | -------- | ----- | -------- | -------- | ------ |"),deepMdTableByRequestParameter(s,e,1)):(markdownLines(e),e.push("None"))}function createApiResponseStatus(t,e){utils.checkUndefined(t.responseCodes)&&t.responseCodes.length>0&&(markdownLines(e),e.push("**status**:"),markdownLines(e),e.push("| code | description | schema |"),e.push("| -------- | -------- | ----- | "),t.responseCodes.forEach(function(s){e.push("|"+utils.toString(s.code,"")+"|"+utils.toString(s.description,"")+"|"+utils.toString(s.schema,"")+"|")}))}function createApiResponseParameters(t,e){if(t.multipartResponseSchema){var s=t.multipCodeDatas;utils.arrNotEmpty(s)&&s.forEach(function(n){markdownLines(e),e.push("**code-"+utils.toString(n.code,"")+"**:"),createApiResponseSingleParam(n,e)})}else createApiResponseSingleParam(t.multipData,e)}function createApiResponseSingleParam(t,e){createApiResponseHeaderParams(t.responseHeaderParameters,e),markdownLines(e),e.push("**Responses**:"),markdownLines(e),utils.arrNotEmpty(t.data)?(e.push("| name | description | type | schema |"),e.push("| -------- | -------- | ----- |----- | "),t.data.forEach(function(s){s.level=1,e.push("|"+getMdTableByLevel(s)+"|"+utils.toString(s.description,"")+"|"+utils.toString(s.type,"")+"|"+utils.toString(s.schemaValue,"")+"|"),deepMdTableByResponseParameter(s.children,e,s.level+1)})):e.push("None"),markdownLines(e),e.push("**Response Sample**:"),t.responseBasicType?(e.push("```text"),e.push(t.responseText),e.push("```")):(e.push("```javascript"),e.push(t.responseValue),e.push("```"))}function deepMdTableByResponseParameter(t,e,s){t!=null&&t!=null&&t.length>0&&t.forEach(function(n){n.level=s,e.push("|"+getMdTableByLevel(n)+"|"+utils.toString(n.description,"")+"|"+utils.toString(n.type,"")+"|"+utils.toString(n.schemaValue,"")+"|"),deepMdTableByResponseParameter(n.children,e,n.level+1)})}function createApiResponseHeaderParams(t,e){utils.checkUndefined(t)&&t.length>0&&(markdownLines(e),e.push("**Response Header**:"),markdownLines(e),e.push("|name | description | type |"),e.push("| -------- | -------- | ----- | "),t.forEach(function(s){e.push("|"+utils.toString(s.name,"")+"|"+utils.toString(s.description,"")+"|"+utils.toString(s.type,"")+"|")}))}function deepMdTableByRequestParameter(t,e,s){t!=null&&t!=null&&t.length>0&&t.forEach(function(n){n.level=s,e.push("|"+getMdTableByLevel(n)+"|"+utils.toString(n.description,"")+"|"+utils.toString(n.in,"")+"|"+utils.toString(n.require,"")+"|"+utils.toString(n.type,"")+"|"+utils.toString(n.schemaValue,"")+"|"),deepMdTableByRequestParameter(n.children,e,n.level+1)})}function getMdTableByLevel(t){for(var e=[],s=1;s<t.level;s++)e.push("&emsp;&emsp;");var n=e.join("")+t.name;return n}const Document_vue_vue_type_style_index_0_scoped_531278ae_lang="",_sfc_main={name:"Document",components:{editor:VAceEditor,DataType:defineAsyncComponent(()=>__vitePreload(()=>import("./DataType-75814406.js"),["./DataType-75814406.js","./doc-7814a93f.js","..\\css\\doc-e469198e.css"],import.meta.url)),EditorShow:defineAsyncComponent(()=>__vitePreload(()=>import("./EditorShow-39bb991a.js"),["./EditorShow-39bb991a.js","./doc-7814a93f.js","..\\css\\doc-e469198e.css","./ext-language_tools-602acc1a.js"],import.meta.url))},props:{api:{type:Object,required:!0},swaggerInstance:{type:Object,required:!0}},setup(){const t=useGlobalsStore(),e=computed(()=>t.language),s=computed(()=>t.swagger),n=computed(()=>t.enableResponseCode),a=useknife4jModels(),{messages:i}=useI18n();return{language:e,swagger:s,responseCodeDisplayStatus:n,knife4jModels:a,messages:i}},data(){return{content:"<span>Hello</span>",contentType:"*/*",columns:[],responseHeaderColumns:[],responseStatuscolumns:[],responseParametersColumns:[],expanRows:!0,multipCode:!1,multipCodeDatas:[],multipData:{},page:!1,reqParameters:[]}},created(){var t=this,e=constants.globalTreeTableModelParams+this.swaggerInstance.id,s=this.swaggerInstance.swaggerTreeTableModels;this.knife4jModels.setValue(e,s),this.initI18n(),this.initRequestParams(),this.initResponseCodeParams(),setTimeout(()=>{t.copyApiAddress(),t.copyApiMarkdown(),t.copyApiUrl()},1500)},watch:{language:function(t,e){this.initI18n()}},methods:{getCurrentI18nInstance(){return this.messages[this.language]},initI18n(){var t=this.getCurrentI18nInstance();this.columns=t.table.documentRequestColumns,this.responseStatuscolumns=t.table.documentResponseStatusColumns,this.responseHeaderColumns=t.table.documentResponseHeaderColumns,this.responseParametersColumns=t.table.documentResponseColumns},copyApiUrl(){var t=this,e="btnCopyMethod"+this.api.id,s=this.api.showUrl,n=new ClipboardJS("#"+e,{text(){return s}});n.on("success",()=>{var a=t.getCurrentI18nInstance(),i=a.message.copy.method.success;message.info(i)}),n.on("error",function(a){var i=t.getCurrentI18nInstance();console.log(i);var r=i.message.copy.method.fail;message.info(r)})},copyApiAddress(){var t=this,e="btnCopyAddress"+this.api.id,s=new ClipboardJS("#"+e,{text(){return window.location.href}});s.on("success",function(n){var a=t.getCurrentI18nInstance(),i=a.message.copy.url.success;message.info(i)}),s.on("error",function(n){var a=t.getCurrentI18nInstance(),i=a.message.copy.url.fail;message.info(i)})},copyApiMarkdown(){var t=this,e="btnCopyMarkdown"+this.api.id,s={...t.api,reqParameters:t.reqParameters,multipCodeDatas:t.multipCodeDatas,multipData:t.multipData},n=new ClipboardJS("#"+e,{text(){var a=t.getCurrentI18nInstance();if(a.lang==="zh")return markdownSingleText(s);if(a.lang==="us")return markdownSingleTextUs(s)}});n.on("success",function(a){var i=t.getCurrentI18nInstance(),r=i.message.copy.document.success;message.info(r)}),n.on("error",function(a){var i=t.getCurrentI18nInstance(),r=i.message.copy.document.fail;message.info(r)})},filterChildrens(t=[],e=[],s){if(t.length===0)return e;const n=this;return(s?e.filter(i=>!t.includes(`${s}.${i.name}`)):e.filter(i=>!t.includes(i.name))).map(i=>(i.id=uniqueId("param"),i.children&&(i.children=n.filterChildrens(t,i.children,i.name)),i))},initRequestParams(){var key=constants.globalTreeTableModelParams+this.swaggerInstance.id,data=[],that=this,apiInfo=this.api;utils.strNotBlank(apiInfo.contentType)&&(this.contentType=apiInfo.contentType),apiInfo.contentType=="application/x-www-form-urlencoded;charset=UTF-8"&&(this.contentType="application/x-www-form-urlencoded");var tmpKeys=Object.keys(apiInfo.ignoreParameters||{}),ignoreParameterAllKeys=[],reg=new RegExp("\\[0\\]","gm");if(tmpKeys!=null&&tmpKeys.length>0&&tmpKeys.forEach(t=>{ignoreParameterAllKeys.push(t),t.indexOf("[0]")>-1&&ignoreParameterAllKeys.push(t.replace(reg,""))}),apiInfo.parameters!=null&&apiInfo.parameters.length>0){var dx=apiInfo.parameters.filter(function(t){return t.name.indexOf("[0]")>-1?ignoreParameterAllKeys.length>0?ignoreParameterAllKeys.filter(e=>!t.name.startsWith(e)).length>0:!0:!ignoreParameterAllKeys.includes(t.name)});data=data.concat(dx)}apiInfo.refTreetableparameters!=null&&apiInfo.refTreetableparameters.length>0&&apiInfo.refTreetableparameters.forEach(function(t){data=data.concat(t.params)}),data!=null&&data.sort(function(t,e){return e.require-t.require});let reqParameters=[];if(data!=null&&data.length>0&&data.forEach(function(param){if(param.pid=="-1"){if(param.children=null,param.schema){var schemaName=param.schemaValue;if(utils.checkUndefined(schemaName)&&that.knife4jModels.exists(key,schemaName)){var model=that.knife4jModels.getByModelName(key,schemaName);if(model=that.swagger.analysisDefinitionRefTableModel(that.swaggerInstance.id,model),model&&model.params){const childrens=model.params.filter(({name})=>!(ignoreParameterAllKeys.includes(name)||ignoreParameterAllKeys.includes(name+"[0]")||ignoreParameterAllKeys.includes(`${param.name}.${name}`)||ignoreParameterAllKeys.some(key=>new RegExp(`^(${key}$|${key}[.[])`).test(name)||eval("/"+key+"/g").test(name)))).map(t=>{const e=that.copyNewParameter(t);if(e.pid=param.id,e.children){const s=JSON.parse(JSON.stringify(e.children)),n=ignoreParameterAllKeys.map(a=>a.startsWith(`${param.name}.${e.name}.`)?a.replace(`${param.name}.${e.name}.`,""):a.startsWith(`${e.name}.`)?a.replace(`${e.name}.`,""):null).filter(Boolean);e.children=that.filterChildrens(n,s)}return e});param.children=childrens.length>0?childrens:null}}}reqParameters.push(param)}}),apiInfo.includeParameters!=null){var tmpIncludeKeys=Object.keys(apiInfo.includeParameters||{}),bodyParam=reqParameters.filter(t=>t.in=="body").length;if(tmpIncludeKeys.length>0&&bodyParam>0){var includeParameters=[],rootKeys=[];this.deepRootKeys(tmpIncludeKeys,rootKeys),reqParameters.forEach(t=>{if(rootKeys.includes(t.name)){var e=cloneDeep(t);e.children=null,t.children!=null&&t.children.length>0&&(e.children=new Array,this.deepIncludeParam(e.name,e,t.children,tmpIncludeKeys,rootKeys)),includeParameters.push(e)}else if(tmpIncludeKeys.includes(t.name)){var e=cloneDeep(t);e.children=null,t.children!=null&&t.children.length>0&&(e.children=new Array,this.deepIncludeParam(e.name,e,t.children,tmpIncludeKeys,rootKeys)),includeParameters.push(e)}}),that.reqParameters=includeParameters}else that.reqParameters=reqParameters}else that.reqParameters=reqParameters},deepRootKeys(t,e){var s=[];t.forEach(n=>{var a=n.substring(0,n.lastIndexOf("."));a.indexOf(".")>-1&&s.push(a),e.includes(a)||e.push(a)}),s.length>0&&this.deepRootKeys(s,e)},deepIncludeParam(t,e,s,n,a){s!=null&&s.length>0&&s.forEach(i=>{var r=t+"."+i.name;if(a.includes(r)){var o=cloneDeep(i);o.children=null,e.children.push(o),utils.arrNotEmpty(i.children)&&(o.children=new Array,this.deepIncludeParam(r,o,i.children,n,a))}else n.includes(r)&&e.children.push(i)})},copyNewParameter(t){const e=n=>n?n.map(a=>{a.id=uniqueId("param"),e(a.children)}):null,s=Object.assign({},t);return s.id=uniqueId("param"),e(s.children),s},deepTreeTableSchemaModel(t,e,s){var n=this;if(utils.checkUndefined(t.schemaValue)){var a=e[t.schemaValue];utils.checkUndefined(a)&&(s.parentTypes.push(t.schemaValue),utils.arrNotEmpty(a.params)&&a.params.forEach(function(i){var r={childrenTypes:i.childrenTypes,def:i.def,description:i.description,enum:i.enum,example:i.example,id:i.id,ignoreFilterName:i.ignoreFilterName,in:i.in,level:i.level,name:i.name,parentTypes:i.parentTypes,pid:i.pid,readOnly:i.readOnly,require:i.require,schema:i.schema,schemaValue:i.schemaValue,show:i.show,txtValue:i.txtValue,type:i.type,validateInstance:i.validateInstance,validateStatus:i.validateStatus,value:i.value};r.pid=t.id,t.children.push(r),r.schema&&s.parentTypes.indexOf(r.schemaValue)==-1&&(r.children=[],n.deepTreeTableSchemaModel(r,e,s),r.children.length==0&&(r.children=null))}))}},findModelChildren(t,e){var s=this;e!=null&&e!=null&&e.length>0&&e.forEach(function(n){n.pid==t.id&&(n.children=[],s.findModelChildren(n,e),n.children.length==0&&(n.children=null),t.children.push(n))})},initResponseCodeParams(){var t=this,e=constants.globalTreeTableModelParams+this.swaggerInstance.id;t.multipCode=this.api.multipartResponseSchema,t.multipCodeDatas=[],t.multipData={};let s=this.api.responseCodes;if(s!=null&&s!=null){s.forEach(function(r){if(r.schema!=null&&r.schema!=null){var o=[];r.responseParameters!=null&&r.responseParameters.length>0&&(o=o.concat(r.responseParameters)),r.responseTreetableRefParameters!=null&&r.responseTreetableRefParameters.length>0&&r.responseTreetableRefParameters.forEach(function(l){o=o.concat(l.params)});let h=[];o!=null&&o.length>0&&o.forEach(function(l){if(l.pid=="-1"){if(l.children=[],l.schema){var p=l.schemaValue;if(utils.checkUndefined(p)&&t.knife4jModels.exists(e,p)){var m=t.knife4jModels.getByModelName(e,p);m=t.swagger.analysisDefinitionRefTableModel(t.swaggerInstance.id,m),m&&m.params&&(l.children=m.params.map(u=>{const c=t.copyNewParameter(u);return c.pid=l.id,c}))}}l.children.length==0&&(l.children=null),h.push(l)}});var d={...r,data:h};t.multipCode||(t.multipData=d),t.multipCodeDatas.push(d)}});var n=Object.keys(t.multipData);if(utils.arrNotEmpty(s)&&!utils.arrNotEmpty(n)){var a=s[0];if(utils.strNotBlank(a.schemaTitle)){var i={...a,data:[]};t.multipData=i}}}},showResponseEditFieldDescription(t){var e=this;if(this.api.multipartResponseSchema){var s=this.multipCodeDatas[0],n="knife4jDocumentShowEditor"+e.api.id+s.code;e.showEditorFieldAnyWay(n)}else{var n="knife4jDocumentShowEditor"+this.api.id;this.showEditorFieldAnyWay(n)}},multipartTabCodeChanges(t){var e=this;setTimeout(()=>{var s="knife4jDocumentShowEditor"+e.api.id+t;e.showEditorFieldAnyWay(s)},1e3)},showEditorFieldAnyWay(t){var e=this.swaggerInstance,s=this.api.getHttpSuccessCodeObject(),n=document.getElementById(t),a=[],i=n.getElementsByClassName("ace_text-layer"),r=n.querySelector(".ace_print-margin").style.left;if(i.length>0)for(var o=i[0].getElementsByClassName("ace_line"),d=0;d<o.length;d++){var h=o[d],l=h.getElementsByClassName("ace_variable"),p=null;if(utils.arrNotEmpty(l)){p=utils.toString(l[0].innerHTML,"").replace(/^"(.*)"$/g,"$1");var m=h.getElementsByClassName("knife4j-debug-editor-field-description");if(!utils.arrNotEmpty(m)&&s!=null){var u=document.createElement("span");u.className="knife4j-debug-editor-field-description",u.innerHTML=s.responseDescriptionFind(a,p,e),u.style.left=r,h.appendChild(u)}}var c=h.getElementsByClassName("ace_paren");if(utils.arrNotEmpty(c)){for(var g=[],f=0;f<c.length;f++)g.push(c[f].innerHTML);var _=g.join("");switch(_){case"[":case"{":a.push(p||0);break;case"}":case"]":a.pop();break}}}}}},_hoisted_1={class:"document"},_hoisted_2={style:{width:"100%"}},_hoisted_3={key:0,class:"knife4j-menu-api-deprecated"},_hoisted_4={key:1},_hoisted_5={class:"knife4j-api-summary"},_hoisted_6={class:"knife4j-api-summary-method"},_hoisted_7={class:"knife4j-api-summary-path"},_hoisted_8={key:0},_hoisted_9=["innerHTML"],_hoisted_10=["innerHTML"],_hoisted_11={key:1},_hoisted_12=["innerHTML"],_hoisted_13=["innerHTML"],_hoisted_14={key:2},_hoisted_15=["innerHTML"],_hoisted_16=["innerHTML"],_hoisted_17=["innerHTML"],_hoisted_18={key:0},_hoisted_19={key:0,style:{color:"red"}},_hoisted_20={key:1},_hoisted_21={key:3},_hoisted_22=["innerHTML"],_hoisted_23={slot:"descriptionTemplate","slot-scope":"text"},_hoisted_24=["innerHTML"],_hoisted_25={slot:"schemaTemplate","slot-scope":"text,record"},_hoisted_26=["innerHTML"],_hoisted_27=["innerHTML"],_hoisted_28={key:4},_hoisted_29={key:0},_hoisted_30=["innerHTML"],_hoisted_31={slot:"descriptionTemplate","slot-scope":"text"},_hoisted_32=["innerHTML"],_hoisted_33=["innerHTML"],_hoisted_34={key:5},_hoisted_35={key:0},_hoisted_36=["innerHTML"],_hoisted_37=["innerHTML"],_hoisted_38={slot:"descriptionTemplate","slot-scope":"text"},_hoisted_39=["innerHTML"],_hoisted_40=["innerHTML"];function _sfc_render(t,e,s,n,a,i){const r=__unplugin_components_5,o=__unplugin_components_1,d=resolveComponent("editor-show"),h=resolveComponent("data-type"),l=__unplugin_components_5$1,p=__unplugin_components_3,m=Tabs;return openBlock(),createElementBlock("div",_hoisted_1,[createBaseVNode("div",_hoisted_2,[createVNode(o,{type:"flex",class:"knife4j-api-title"},{default:withCtx(()=>[createVNode(r,{flex:18},{default:withCtx(()=>[s.api.deprecated?(openBlock(),createElementBlock("span",_hoisted_3,toDisplayString(s.api.summary),1)):(openBlock(),createElementBlock("span",_hoisted_4,toDisplayString(s.api.summary),1))]),_:1}),createVNode(r,{flex:2,id:"btnCopyMethod"+s.api.id,class:"knife4j-api-copy-address",innerHTML:t.$t("doc.copyMethod")},null,8,["id","innerHTML"]),createVNode(r,{flex:2,id:"btnCopyMarkdown"+s.api.id,class:"knife4j-api-copy-address",innerHTML:t.$t("doc.copy")},null,8,["id","innerHTML"]),createVNode(r,{flex:2,id:"btnCopyAddress"+s.api.id,class:"knife4j-api-copy-address",innerHTML:t.$t("doc.copyHash")},null,8,["id","innerHTML"])]),_:1}),createVNode(o,{class:normalizeClass("knife4j-api-"+s.api.methodType.toLowerCase())},{default:withCtx(()=>[createBaseVNode("div",_hoisted_5,[createBaseVNode("span",_hoisted_6,toDisplayString(s.api.methodType),1),createBaseVNode("span",_hoisted_7,toDisplayString(s.api.showUrl),1)])]),_:1},8,["class"]),createVNode(o,{class:"knife4j-api-row"},{default:withCtx(()=>[createVNode(r,{span:12},{default:withCtx(()=>[createVNode(o,null,{default:withCtx(()=>[createVNode(r,{class:"api-basic-title",span:6,innerHTML:t.$t("doc.produces")},null,8,["innerHTML"]),createTextVNode(" "+toDisplayString(a.contentType),1)]),_:1})]),_:1}),createVNode(r,{span:12},{default:withCtx(()=>[createVNode(o,null,{default:withCtx(()=>[createVNode(r,{class:"api-basic-title",span:6,innerHTML:t.$t("doc.consumes")},null,8,["innerHTML"]),createTextVNode(" "+toDisplayString(s.api.produces),1)]),_:1})]),_:1})]),_:1})]),s.api.author?(openBlock(),createElementBlock("div",_hoisted_8,[createBaseVNode("div",{class:"api-title",innerHTML:t.$t("doc.author")},null,8,_hoisted_9),s.api.author?(openBlock(),createElementBlock("div",{key:0,innerHTML:s.api.author,class:"api-body-desc"},null,8,_hoisted_10)):createCommentVNode("",!0)])):createCommentVNode("",!0),s.api.description?(openBlock(),createElementBlock("div",_hoisted_11,[createBaseVNode("div",{class:"api-title",innerHTML:t.$t("doc.des")},null,8,_hoisted_12),s.api.description?(openBlock(),createElementBlock("div",{key:0,innerHTML:s.api.description,class:"api-body-desc"},null,8,_hoisted_13)):createCommentVNode("",!0)])):createCommentVNode("",!0),s.api.requestValue?(openBlock(),createElementBlock("div",_hoisted_14,[createBaseVNode("div",{class:"api-title",innerHTML:t.$t("doc.requestExample")},null,8,_hoisted_15),createVNode(d,{value:s.api.requestValue,xmlMode:s.api.xmlRequest},null,8,["value","xmlMode"])])):createCommentVNode("",!0),createBaseVNode("div",{class:"api-title",innerHTML:t.$t("doc.params")},null,8,_hoisted_16),createVNode(l,{defaultExpandAllRows:a.expanRows,columns:a.columns,dataSource:a.reqParameters,rowKey:"id",size:"small",pagination:a.page},{bodyCell:withCtx(({column:u,record:c})=>[u.dataIndex==="description"?(openBlock(),createElementBlock(Fragment,{key:0},[createBaseVNode("span",{innerHTML:c.description},null,8,_hoisted_17),c.example?(openBlock(),createElementBlock("span",_hoisted_18,",示例值("+toDisplayString(c.example)+")",1)):createCommentVNode("",!0)],64)):u.dataIndex==="in"?(openBlock(),createElementBlock("span",{key:1,class:normalizeClass("knife4j-request-"+c.in)},toDisplayString(c.in),3)):u.dataIndex==="require"?(openBlock(),createElementBlock(Fragment,{key:2},[c.require?(openBlock(),createElementBlock("span",_hoisted_19,toDisplayString(c.require.toLocaleString()),1)):(openBlock(),createElementBlock("span",_hoisted_20,toDisplayString(c.require.toLocaleString()),1))],64)):u.dataIndex==="type"?(openBlock(),createBlock(h,{key:3,text:c.type,record:c},null,8,["text","record"])):createCommentVNode("",!0)]),_:1},8,["defaultExpandAllRows","columns","dataSource","pagination"]),n.responseCodeDisplayStatus?(openBlock(),createElementBlock("div",_hoisted_21,[createBaseVNode("div",{class:"api-title",innerHTML:t.$t("doc.response")},null,8,_hoisted_22),createVNode(l,{defaultExpandAllRows:a.expanRows,columns:a.responseStatuscolumns,dataSource:s.api.responseCodes,rowKey:"code",size:"small",pagination:a.page},{default:withCtx(()=>[createBaseVNode("template",_hoisted_23,[createBaseVNode("div",{innerHTML:t.text},null,8,_hoisted_24)]),createBaseVNode("template",_hoisted_25,[t.text!=null?(openBlock(),createElementBlock("span",{key:0,innerHTML:t.text},null,8,_hoisted_26)):t.record.schemaTitle!=null?(openBlock(),createElementBlock("span",{key:1,innerHTML:t.record.schemaTitle},null,8,_hoisted_27)):createCommentVNode("",!0)])]),_:1},8,["defaultExpandAllRows","columns","dataSource","pagination"])])):createCommentVNode("",!0),s.api.multipartResponseSchema?(openBlock(),createElementBlock("div",_hoisted_28,[createVNode(m,{onChange:i.multipartTabCodeChanges},{default:withCtx(()=>[(openBlock(!0),createElementBlock(Fragment,null,renderList(a.multipCodeDatas,u=>(openBlock(),createBlock(p,{key:u.code,tab:t.$t("doc.responseHeaderParams")},{default:withCtx(()=>[u.responseHeaderParameters?(openBlock(),createElementBlock("div",_hoisted_29,[createVNode(l,{defaultExpandAllRows:a.expanRows,columns:a.responseHeaderColumns,dataSource:u.responseHeaderParameters,rowKey:"id",size:"small",pagination:a.page},null,8,["defaultExpandAllRows","columns","dataSource","pagination"])])):createCommentVNode("",!0),createBaseVNode("div",{class:"api-title",innerHTML:t.$t("doc.responseParams")},null,8,_hoisted_30),createVNode(l,{defaultExpandAllRows:a.expanRows,columns:a.responseParametersColumns,dataSource:u.data,rowKey:"id",size:"small",pagination:a.page},{default:withCtx(()=>[createBaseVNode("template",_hoisted_31,[createBaseVNode("span",{innerHTML:t.text},null,8,_hoisted_32)])]),_:2},1032,["defaultExpandAllRows","columns","dataSource","pagination"]),createBaseVNode("div",{class:"api-title",innerHTML:t.$t("doc.responseExample")},null,8,_hoisted_33),createVNode(o,{id:"knife4jDocumentShowEditor"+s.api.id+u.code,style:{display:"block"}},{default:withCtx(()=>[createVNode(d,{onShowDescription:i.showResponseEditFieldDescription,value:u.responseBasicType?u.responseText:u.responseValue},null,8,["onShowDescription","value"])]),_:2},1032,["id"])]),_:2},1032,["tab"]))),128))]),_:1},8,["onChange"])])):(openBlock(),createElementBlock("div",_hoisted_34,[s.api.responseHeaderParameters?(openBlock(),createElementBlock("div",_hoisted_35,[createBaseVNode("div",{class:"api-title",innerHTML:t.$t("doc.responseHeaderParams")},null,8,_hoisted_36),createVNode(l,{defaultExpandAllRows:a.expanRows,columns:a.responseHeaderColumns,dataSource:s.api.responseHeaderParameters,rowKey:"id",size:"small",pagination:a.page},null,8,["defaultExpandAllRows","columns","dataSource","pagination"])])):createCommentVNode("",!0),createBaseVNode("div",{class:"api-title",innerHTML:t.$t("doc.responseParams")},null,8,_hoisted_37),createVNode(l,{defaultExpandAllRows:a.expanRows,columns:a.responseParametersColumns,dataSource:a.multipData.data,rowKey:"id",size:"small",pagination:a.page},{default:withCtx(()=>[createBaseVNode("template",_hoisted_38,[createBaseVNode("span",{innerHTML:t.text},null,8,_hoisted_39)])]),_:1},8,["defaultExpandAllRows","columns","dataSource","pagination"]),createBaseVNode("div",{class:"api-title",innerHTML:t.$t("doc.responseExample")},null,8,_hoisted_40),createVNode(o,{id:"knife4jDocumentShowEditor"+s.api.id,style:{display:"block"}},{default:withCtx(()=>[createVNode(d,{onShowDescription:i.showResponseEditFieldDescription,value:a.multipData.responseBasicType?a.multipData.responseText:a.multipData.responseValue},null,8,["onShowDescription","value"])]),_:1},8,["id"])]))])}const Document=_export_sfc(_sfc_main,[["render",_sfc_render],["__scopeId","data-v-531278ae"]]);export{Document as default};
