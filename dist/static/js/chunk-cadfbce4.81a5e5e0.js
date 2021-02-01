(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-cadfbce4"],{"22ce":function(e,t,r){"use strict";r.d(t,"f",(function(){return i})),r.d(t,"d",(function(){return l})),r.d(t,"e",(function(){return o})),r.d(t,"g",(function(){return s})),r.d(t,"b",(function(){return n})),r.d(t,"i",(function(){return c})),r.d(t,"c",(function(){return u})),r.d(t,"j",(function(){return m})),r.d(t,"a",(function(){return d})),r.d(t,"h",(function(){return _}));var a=r("b775");function i(e){return Object(a["a"])({url:"/service/service_list",method:"get",params:e})}function l(e){return Object(a["a"])({url:"/service/service_delete",method:"get",params:e})}function o(e){return Object(a["a"])({url:"/service/service_detail",method:"get",params:e})}function s(e){return Object(a["a"])({url:"/service/service_stat",method:"get",params:e})}function n(e){return Object(a["a"])({url:"/service/service_add_http",method:"post",data:e})}function c(e){return Object(a["a"])({url:"/service/service_update_http",method:"post",data:e})}function u(e){return Object(a["a"])({url:"/service/service_add_tcp",method:"post",data:e})}function m(e){return Object(a["a"])({url:"/service/service_update_tcp",method:"post",data:e})}function d(e){return Object(a["a"])({url:"/service/service_add_grpc",method:"post",data:e})}function _(e){return Object(a["a"])({url:"/service/service_update_grpc",method:"post",data:e})}},"2d0b":function(e,t,r){},7207:function(e,t,r){"use strict";r.r(t);var a=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"mixin-components-container"},[r("el-row",[r("el-card",{staticClass:"box-card"},[r("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[!1===e.isEdit?r("span",[e._v("Create HTTP Service")]):e._e(),!0===e.isEdit?r("span",[e._v("Update HTTP Service")]):e._e()]),r("div",{staticStyle:{"margin-bottom":"50px"}},[r("el-form",{ref:"form",attrs:{model:e.form,"label-width":"140px"}},[r("el-form-item",{staticClass:"is-required",attrs:{label:"Service Name"}},[r("el-input",{attrs:{placeholder:"6-128 characters",disabled:!0===e.isEdit},model:{value:e.form.service_name,callback:function(t){e.$set(e.form,"service_name",t)},expression:"form.service_name"}})],1),r("el-form-item",{staticClass:"is-required",attrs:{label:" Description"}},[r("el-input",{attrs:{placeholder:"Maximum 255 characters"},model:{value:e.form.service_desc,callback:function(t){e.$set(e.form,"service_desc",t)},expression:"form.service_desc"}})],1),r("el-form-item",{staticClass:"is-required",attrs:{label:"Access Type"}},[r("el-input",{staticClass:"input-with-select",attrs:{disabled:!0===e.isEdit,placeholder:"Path：/user/, Domain：www.test.com"},model:{value:e.form.rule,callback:function(t){e.$set(e.form,"rule",t)},expression:"form.rule"}},[r("el-select",{staticStyle:{width:"95px"},attrs:{slot:"prepend",placeholder:"Select",disabled:!0===e.isEdit},slot:"prepend",model:{value:e.form.rule_type,callback:function(t){e.$set(e.form,"rule_type",t)},expression:"form.rule_type"}},[r("el-option",{attrs:{label:"Path",value:0}}),r("el-option",{attrs:{label:"Domain",value:1}})],1)],1)],1),r("el-form-item",{attrs:{label:"https"}},[r("el-switch",{attrs:{"active-value":1,"inactive-value":0},model:{value:e.form.need_https,callback:function(t){e.$set(e.form,"need_https",t)},expression:"form.need_https"}}),r("span",{staticStyle:{color:"#606266","font-weight":"700"}},[e._v(" strip_uri ")]),r("el-switch",{attrs:{"active-value":1,"inactive-value":0},model:{value:e.form.need_strip_uri,callback:function(t){e.$set(e.form,"need_strip_uri",t)},expression:"form.need_strip_uri"}})],1),r("el-form-item",{attrs:{label:"URL Rewrite"}},[r("el-input",{attrs:{type:"textarea",autosize:"",placeholder:"Format：^/gateway/test_service(.*) $1"},model:{value:e.form.url_rewrite,callback:function(t){e.$set(e.form,"url_rewrite",t)},expression:"form.url_rewrite"}})],1),r("el-form-item",{attrs:{label:"Header Transfer"}},[r("el-input",{attrs:{type:"textarea",autosize:"",placeholder:"add/del/edit eg. add headerName headValue"},model:{value:e.form.header_transfer,callback:function(t){e.$set(e.form,"header_transfer",t)},expression:"form.header_transfer"}})],1),r("el-form-item",{attrs:{label:"Authentication"}},[r("el-switch",{attrs:{"active-value":1,"inactive-value":0},model:{value:e.form.open_auth,callback:function(t){e.$set(e.form,"open_auth",t)},expression:"form.open_auth"}})],1),r("el-form-item",{attrs:{label:"IP Whitelist"}},[r("el-input",{attrs:{type:"textarea",autosize:"",placeholder:"eg. 127.0.0.1 multiple lines"},model:{value:e.form.white_list,callback:function(t){e.$set(e.form,"white_list",t)},expression:"form.white_list"}})],1),r("el-form-item",{attrs:{label:"IP Blacklist"}},[r("el-input",{attrs:{type:"textarea",autosize:"",placeholder:"eg. 127.0.0.1 multiple lines"},model:{value:e.form.black_list,callback:function(t){e.$set(e.form,"black_list",t)},expression:"form.black_list"}})],1),r("el-form-item",{attrs:{label:"Client limit"}},[r("el-input",{attrs:{placeholder:"0 means no limits"},model:{value:e.form.clientip_flow_limit,callback:function(t){e.$set(e.form,"clientip_flow_limit",t)},expression:"form.clientip_flow_limit"}})],1),r("el-form-item",{attrs:{label:"Server limit"}},[r("el-input",{attrs:{placeholder:"0 means no limits"},model:{value:e.form.service_flow_limit,callback:function(t){e.$set(e.form,"service_flow_limit",t)},expression:"form.service_flow_limit"}})],1),r("el-form-item",{attrs:{label:"Round type"}},[r("el-radio-group",{model:{value:e.form.round_type,callback:function(t){e.$set(e.form,"round_type",t)},expression:"form.round_type"}},[r("el-radio",{attrs:{label:"0"},model:{value:e.form.round_type,callback:function(t){e.$set(e.form,"round_type",t)},expression:"form.round_type"}},[e._v("random")]),r("el-radio",{attrs:{label:"1"},model:{value:e.form.round_type,callback:function(t){e.$set(e.form,"round_type",t)},expression:"form.round_type"}},[e._v("round-robin")]),r("el-radio",{attrs:{label:"2"},model:{value:e.form.round_type,callback:function(t){e.$set(e.form,"round_type",t)},expression:"form.round_type"}},[e._v("weight_round-robin")]),r("el-radio",{attrs:{label:"3"},model:{value:e.form.round_type,callback:function(t){e.$set(e.form,"round_type",t)},expression:"form.round_type"}},[e._v("ip_hash")])],1)],1),r("el-form-item",{staticClass:"is-required",attrs:{label:"IP list"}},[r("el-input",{attrs:{type:"textarea",autosize:"",placeholder:"eg. 127.0.0.1 multiple lines"},model:{value:e.form.ip_list,callback:function(t){e.$set(e.form,"ip_list",t)},expression:"form.ip_list"}})],1),r("el-form-item",{staticClass:"is-required",attrs:{label:"Weight list"}},[r("el-input",{attrs:{type:"textarea",autosize:"",placeholder:"eg. 127.0.0.1 multiple lines"},model:{value:e.form.weight_list,callback:function(t){e.$set(e.form,"weight_list",t)},expression:"form.weight_list"}})],1),r("el-form-item",{attrs:{label:"Connect timeout"}},[r("el-input",{attrs:{placeholder:"0 means no limits"},model:{value:e.form.upstream_connect_timeout,callback:function(t){e.$set(e.form,"upstream_connect_timeout",t)},expression:"form.upstream_connect_timeout"}})],1),r("el-form-item",{attrs:{label:"Header timeout"}},[r("el-input",{attrs:{placeholder:"0 means no limits"},model:{value:e.form.upstream_header_timeout,callback:function(t){e.$set(e.form,"upstream_header_timeout",t)},expression:"form.upstream_header_timeout"}})],1),r("el-form-item",{attrs:{label:"Idle timeout"}},[r("el-input",{attrs:{placeholder:"0 means no limits"},model:{value:e.form.upstream_idle_timeout,callback:function(t){e.$set(e.form,"upstream_idle_timeout",t)},expression:"form.upstream_idle_timeout"}})],1),r("el-form-item",{attrs:{label:"Idle connection"}},[r("el-input",{attrs:{placeholder:"0 means no limits"},model:{value:e.form.upstream_max_idle,callback:function(t){e.$set(e.form,"upstream_max_idle",t)},expression:"form.upstream_max_idle"}})],1),r("el-form-item",[r("el-button",{attrs:{type:"primary",disabled:e.submitButDisabled},on:{click:e.handleSubmit}},[e._v("Submit")])],1)],1)],1)])],1)],1)},i=[],l=(r("a9e3"),r("d3b7"),r("ac1f"),r("25f0"),r("5319"),r("22ce")),o={name:"ServiceCreateHttp",data:function(){return{isEdit:!1,submitButDisabled:!1,form:{service_name:"",service_desc:"",rule_type:0,rule:"",need_https:0,need_websocket:0,need_strip_uri:1,url_rewrite:"",header_transfer:"",round_type:2,ip_list:"",weight_list:"",upstream_connect_timeout:"",upstream_header_timeout:"",upstream_idle_timeout:"",upstream_max_idle:"",open_auth:0,black_list:"",white_list:"",clientip_flow_limit:"",service_flow_limit:""}}},created:function(){var e=this.$route.params&&this.$route.params.id;e>0&&(this.isEdit=!0,this.fetchData(e))},methods:{fetchData:function(e){var t=this;console.log("fetch by ID: "+e);var r={id:e};Object(l["e"])(r).then((function(e){t.form.id=e.data.info.id,t.form.load_type=e.data.info.load_type,t.form.service_name=e.data.info.service_name,t.form.service_desc=e.data.info.service_desc,t.form.rule_type=e.data.http_rule.rule_type,t.form.rule=e.data.http_rule.rule,t.form.need_https=e.data.http_rule.need_https,t.form.need_websocket=e.data.http_rule.need_websocket,t.form.need_strip_uri=e.data.http_rule.need_strip_uri,t.form.url_rewrite=e.data.http_rule.url_rewrite.replace(/,/g,"\n"),t.form.header_transfer=e.data.http_rule.header_transfer.replace(/,/g,"\n"),t.form.round_type=e.data.load_balance.round_type.toString(),t.form.ip_list=e.data.load_balance.ip_list.replace(/,/g,"\n"),t.form.weight_list=e.data.load_balance.weight_list.replace(/,/g,"\n"),t.form.upstream_connect_timeout=e.data.load_balance.upstream_connect_timeout,t.form.upstream_header_timeout=e.data.load_balance.upstream_header_timeout,t.form.upstream_idle_timeout=e.data.load_balance.upstream_idle_timeout,t.form.upstream_max_idle=e.data.load_balance.upstream_max_idle,t.form.upstream_max_idle=e.data.load_balance.upstream_max_idle,t.form.open_auth=e.data.access_control.open_auth,t.form.black_list=e.data.access_control.black_list.replace(/,/g,"\n"),t.form.white_list=e.data.access_control.white_list.replace(/,/g,"\n"),t.form.clientip_flow_limit=e.data.access_control.clientip_flow_limit,t.form.service_flow_limit=e.data.access_control.service_flow_limit,console.log(e)})).catch((function(e){console.log("fetch data error: "+e)}))},handleSubmit:function(){var e=this;this.submitButDisabled=!0;var t=Object.assign({},this.form);console.log(t),t.url_rewrite=t.url_rewrite.replace(/\n/g,","),t.header_transfer=t.header_transfer.replace(/\n/g,","),t.round_type=Number(t.round_type),t.ip_list=t.ip_list.replace(/\n/g,","),t.weight_list=t.weight_list.replace(/\n/g,","),t.black_list=t.black_list.replace(/\n/g,","),t.white_list=t.white_list.replace(/\n/g,","),t.round_type=Number(t.round_type),t.service_flow_limit=Number(t.service_flow_limit),t.clientip_flow_limit=Number(t.clientip_flow_limit),t.upstream_connect_timeout=Number(t.upstream_connect_timeout),t.upstream_header_timeout=Number(t.upstream_header_timeout),t.upstream_idle_timeout=Number(t.upstream_idle_timeout),t.upstream_max_idle=Number(t.upstream_max_idle),this.isEdit?Object(l["i"])(t).then((function(t){e.submitButDisabled=!1,e.$notify({title:"Success",message:"Updated",type:"success",duration:2e3})})).catch((function(){e.submitButDisabled=!1})):Object(l["b"])(t).then((function(t){e.submitButDisabled=!1,e.$notify({title:"Success",message:"Added",type:"success",duration:2e3})})).catch((function(){e.submitButDisabled=!1}))}}},s=o,n=(r("8d77"),r("2877")),c=Object(n["a"])(s,a,i,!1,null,"52d65716",null);t["default"]=c.exports},"8d77":function(e,t,r){"use strict";r("2d0b")}}]);