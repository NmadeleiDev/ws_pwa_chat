(function(e){function t(t){for(var a,r,i=t[0],c=t[1],u=t[2],l=0,h=[];l<i.length;l++)r=i[l],Object.prototype.hasOwnProperty.call(s,r)&&s[r]&&h.push(s[r][0]),s[r]=0;for(a in c)Object.prototype.hasOwnProperty.call(c,a)&&(e[a]=c[a]);d&&d(t);while(h.length)h.shift()();return o.push.apply(o,u||[]),n()}function n(){for(var e,t=0;t<o.length;t++){for(var n=o[t],a=!0,r=1;r<n.length;r++){var i=n[r];0!==s[i]&&(a=!1)}a&&(o.splice(t--,1),e=c(c.s=n[0]))}return e}var a={},r={app:0},s={app:0},o=[];function i(e){return c.p+"js/"+({about:"about"}[e]||e)+"."+{about:"27df7c89"}[e]+".js"}function c(t){if(a[t])return a[t].exports;var n=a[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,c),n.l=!0,n.exports}c.e=function(e){var t=[],n={about:1};r[e]?t.push(r[e]):0!==r[e]&&n[e]&&t.push(r[e]=new Promise((function(t,n){for(var a="css/"+({about:"about"}[e]||e)+"."+{about:"296c7e24"}[e]+".css",s=c.p+a,o=document.getElementsByTagName("link"),i=0;i<o.length;i++){var u=o[i],l=u.getAttribute("data-href")||u.getAttribute("href");if("stylesheet"===u.rel&&(l===a||l===s))return t()}var h=document.getElementsByTagName("style");for(i=0;i<h.length;i++){u=h[i],l=u.getAttribute("data-href");if(l===a||l===s)return t()}var d=document.createElement("link");d.rel="stylesheet",d.type="text/css",d.onload=t,d.onerror=function(t){var a=t&&t.target&&t.target.src||s,o=new Error("Loading CSS chunk "+e+" failed.\n("+a+")");o.code="CSS_CHUNK_LOAD_FAILED",o.request=a,delete r[e],d.parentNode.removeChild(d),n(o)},d.href=s;var p=document.getElementsByTagName("head")[0];p.appendChild(d)})).then((function(){r[e]=0})));var a=s[e];if(0!==a)if(a)t.push(a[2]);else{var o=new Promise((function(t,n){a=s[e]=[t,n]}));t.push(a[2]=o);var u,l=document.createElement("script");l.charset="utf-8",l.timeout=120,c.nc&&l.setAttribute("nonce",c.nc),l.src=i(e);var h=new Error;u=function(t){l.onerror=l.onload=null,clearTimeout(d);var n=s[e];if(0!==n){if(n){var a=t&&("load"===t.type?"missing":t.type),r=t&&t.target&&t.target.src;h.message="Loading chunk "+e+" failed.\n("+a+": "+r+")",h.name="ChunkLoadError",h.type=a,h.request=r,n[1](h)}s[e]=void 0}};var d=setTimeout((function(){u({type:"timeout",target:l})}),12e4);l.onerror=l.onload=u,document.head.appendChild(l)}return Promise.all(t)},c.m=e,c.c=a,c.d=function(e,t,n){c.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},c.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},c.t=function(e,t){if(1&t&&(e=c(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(c.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var a in e)c.d(n,a,function(t){return e[t]}.bind(null,a));return n},c.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return c.d(t,"a",t),t},c.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},c.p="/",c.oe=function(e){throw console.error(e),e};var u=window["webpackJsonp"]=window["webpackJsonp"]||[],l=u.push.bind(u);u.push=t,u=u.slice();for(var h=0;h<u.length;h++)t(u[h]);var d=l;o.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("cd49")},"034f":function(e,t,n){"use strict";var a=n("8a23"),r=n.n(a);r.a},"8a23":function(e,t,n){},cd49:function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var a=n("2b0e"),r=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-app",[n("router-view")],1)},s=[],o=a["a"].extend({name:"App",components:{},data:function(){return{}}}),i=o,c=(n("034f"),n("2877")),u=n("6544"),l=n.n(u),h=n("7496"),d=Object(c["a"])(i,r,s,!1,null,null,null),p=d.exports;l()(d,{VApp:h["a"]});var f=n("9483");Object(f["a"])("".concat("/","service-worker.js"),{ready:function(){console.log("App is being served from cache by a service worker.\nFor more details, visit https://goo.gl/AFskqB")},registered:function(){console.log("Service worker has been registered.")},cached:function(){console.log("Content has been cached for offline use.")},updatefound:function(){console.log("New content is downloading.")},updated:function(){console.log("New content is available; please refresh.")},offline:function(){console.log("No internet connection found. App is running in offline mode.")},error:function(e){console.error("Error during service worker registration:",e)}});n("d3b7");var g=n("8c4f"),m=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"default-dark-db fill-height"},[n("LoginAppBar"),n("v-main",{staticClass:"default-dark-db fill-height"},[n("v-row",{staticClass:"fill-height",attrs:{"align-content":"center"}},[n("v-col",{staticClass:"d-flex fill-height flex-column justify-space-between align-center",attrs:{"align-self":"center"}},[n("v-card",{staticClass:"w-auto h-100 default-dark-db",attrs:{color:"#f6f6f6",elevation:"0"}},[n("v-card-title",{staticClass:"d-flex flex-row justify-center"},[e._v(" Welcome back! ")]),n("v-card-text",[n("transition",{attrs:{"enter-active-class":"animate__bounceInLeft"}},[e.shownFields?n("v-text-field",{attrs:{placeholder:"Login",color:"green"},model:{value:e.login,callback:function(t){e.login=t},expression:"login"}}):e._e()],1),n("transition",{attrs:{"enter-active-class":"animate__bounceInRight"}},[e.shownFields?n("v-text-field",{attrs:{placeholder:"Password",color:"green",type:"password"},model:{value:e.password,callback:function(t){e.password=t},expression:"password"}}):e._e()],1),n("transition",{attrs:{"enter-active-class":"animate__bounceInLeft"}},[!e.isLogin&&e.shownFields?n("v-text-field",{attrs:{placeholder:"Verify password",type:"password",color:"green"},model:{value:e.verifyPassword,callback:function(t){e.verifyPassword=t},expression:"verifyPassword"}}):e._e()],1)],1),n("v-card-actions",{staticClass:"d-flex flex-row justify-center"},[n("v-btn",{attrs:{large:"",disabled:!e.isInputValid,color:e.isInputValid?"green":"black"},on:{click:function(t){return e.passThrough()}}},[e._v(" "+e._s(e.isLogin?"Sign in":"Sign up")+" ")])],1)],1),n("v-banner",{staticClass:"mb-10"},[e._v(" "+e._s(e.isLogin?"Don't have an account yet?":"Already have an account?")+" "),n("v-btn",{attrs:{text:""},on:{click:function(t){e.isLogin=!e.isLogin}}},[e._v(" "+e._s(e.isLogin?"Sign up":"Sign in")+" ")])],1)],1)],1)],1)],1)},v=[],b=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-app-bar",{attrs:{color:"accent-4",dense:"",dark:""}},[n("v-toolbar-title",[e._v("Login")]),n("v-spacer")],1)},w=[],y=a["a"].extend({name:"LoginAppBar"}),k=y,j=n("40dc"),O=n("2fa4"),S=n("2a7f"),C=Object(c["a"])(k,b,w,!1,null,"19bf3c26",null),x=C.exports;l()(C,{VAppBar:j["a"],VSpacer:O["a"],VToolbarTitle:S["a"]});var L=a["a"].extend({name:"Login",components:{LoginAppBar:x},data:function(){return{login:"",password:"",verifyPassword:"",isLogin:!0,shownFields:!1}},created:function(){this.$store.dispatch("findLocalKeys").then((function(e){!0===e&&K.push("/home")}))},mounted:function(){this.shownFields=!0},methods:{passThrough:function(){this.isLogin?this.$store.dispatch("signIn",{login:this.login,password:this.password}).then((function(e){!0===e&&K.push("/home")})):this.$store.dispatch("signUp",{login:this.login,password:this.password}).then((function(e){!0===e&&K.push("/home")}))}},computed:{isInputValid:function(){return this.isLogin?""!==this.login&&""!==this.password:""!==this.login&&""!==this.password&&this.password===this.verifyPassword}}}),I=L,_=(n("eeb1"),n("e4e5")),A=n("8336"),T=n("b0af"),R=n("99d9"),E=n("62ad"),N=n("f6c4"),F=n("0fd9"),U=n("8654"),V=Object(c["a"])(I,m,v,!1,null,null,null),P=V.exports;l()(V,{VBanner:_["a"],VBtn:A["a"],VCard:T["a"],VCardActions:R["a"],VCardText:R["b"],VCardTitle:R["c"],VCol:E["a"],VMain:N["a"],VRow:F["a"],VTextField:U["a"]}),a["a"].use(g["a"]);var M=[{path:"/",name:"Login",component:P},{path:"/home",name:"Home",component:function(){return n.e("about").then(n.bind(null,"cbe8"))}},{path:"/users",name:"All users",component:function(){return n.e("about").then(n.bind(null,"4ca1"))}},{path:"/chat",name:"Chat here",component:function(){return n.e("about").then(n.bind(null,"5c2b"))}}],D=new g["a"]({routes:M}),K=D,B=n("2f62"),$=n("675f"),q=(n("96cf"),n("1da1")),J=n("bee2"),W=n("262e"),H=n("2caf"),G=n("d4ec"),z=n("bc3a"),Q=n.n(z);function X(e){return"/api/v1/"+e}var Y=function(){var e=Object(q["a"])(regeneratorRuntime.mark((function e(t){var n,a,r;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return n=X(t),e.prev=1,e.next=4,Q.a.get(n,{headers:{mobile:"true"}});case 4:a=e.sent,e.next=11;break;case 7:return e.prev=7,e.t0=e["catch"](1),console.log("Request get error: ",e.t0),e.abrupt("return",{status:!1,data:null});case 11:return r=a.data.status,!0!==r&&console.log(a),e.abrupt("return",a.data);case 14:case"end":return e.stop()}}),e,null,[[1,7]])})));return function(t){return e.apply(this,arguments)}}(),Z=function(){var e=Object(q["a"])(regeneratorRuntime.mark((function e(t,n,a){var r,s;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return r=X(t),e.prev=1,e.next=4,Q.a.post(r,n,{headers:{mobile:"true","Event-date":a}});case 4:s=e.sent,e.next=11;break;case 7:return e.prev=7,e.t0=e["catch"](1),console.log("Request post error: ",e.t0),e.abrupt("return",{status:!1,data:null});case 11:return!0!==s.data.status&&console.log("failed response: ",s),e.abrupt("return",s.data);case 13:case"end":return e.stop()}}),e,null,[[1,7]])})));return function(t,n,a){return e.apply(this,arguments)}}(),ee=function(){var e=Object(q["a"])(regeneratorRuntime.mark((function e(t,n,a){var r,s;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return r=X(t),e.prev=1,e.next=4,Q.a.put(r,n,{headers:{mobile:"true","Event-date":a}});case 4:s=e.sent,e.next=11;break;case 7:return e.prev=7,e.t0=e["catch"](1),console.log("Request put error: ",e.t0),e.abrupt("return",{status:!1,data:null});case 11:return!0!==s.data.status&&console.log("failed response: ",s),e.abrupt("return",s.data);case 13:case"end":return e.stop()}}),e,null,[[1,7]])})));return function(t,n,a){return e.apply(this,arguments)}}(),te={get:Y,put:ee,post:Z},ne=(n("a15b"),n("25f0"),n("6c27")),ae=function(e){return Object(ne["sha256"])([e.username,e.password,e.username].join("&-")).toString()},re=function(e){return Object(ne["sha256"])([e.username,e.timeStamp,e.sessionKey,e.userSecret].join("")).toString()},se={generateToken:re,generateUserSecret:ae},oe=function e(){Object(G["a"])(this,e),this.isLogged=!1},ie=function(e){Object(W["a"])(n,e);var t=Object(H["a"])(n);function n(){return Object(G["a"])(this,n),t.apply(this,arguments)}return Object(J["a"])(n,[{key:"isLogged",get:function(){return this.state.isLogged}}]),n}($["b"]),ce=function(e){Object(W["a"])(n,e);var t=Object(H["a"])(n);function n(){return Object(G["a"])(this,n),t.apply(this,arguments)}return Object(J["a"])(n,[{key:"setLoginState",value:function(e){this.state.isLogged=e}}]),n}($["d"]),ue=function(e){Object(W["a"])(n,e);var t=Object(H["a"])(n);function n(){return Object(G["a"])(this,n),t.apply(this,arguments)}return Object(J["a"])(n,[{key:"signUp",value:function(){var e=Object(q["a"])(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return console.log("Tring to sign up: ",t),e.next=3,te.post("signup",{data:{username:t.login,password:t.password},auth:null},"");case 3:if(n=e.sent,!0!==n.status){e.next=11;break}return localStorage.setItem("username",t.login),localStorage.setItem("sessionKey",n.data.token),localStorage.setItem("userSecret",se.generateUserSecret({username:t.login,password:t.password})),this.commit("setLoginState",!0),console.log("Signed up successfully: ",localStorage.getItem("sessionKey"),localStorage.getItem("userSecret")),e.abrupt("return",!0);case 11:return e.abrupt("return",!1);case 12:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}()},{key:"signIn",value:function(){var e=Object(q["a"])(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,te.post("signin",{data:{username:t.login,password:t.password},auth:null},"");case 2:if(n=e.sent,!0!==n.status){e.next=10;break}return localStorage.setItem("username",t.login),localStorage.setItem("sessionKey",n.data.token),localStorage.setItem("userSecret",se.generateUserSecret({username:t.login,password:t.password})),this.commit("setLoginState",!0),console.log("Signed in successfully: ",localStorage.getItem("sessionKey"),localStorage.getItem("userSecret")),e.abrupt("return",!0);case 10:return e.abrupt("return",!1);case 11:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}()},{key:"findLocalKeys",value:function(){var e=Object(q["a"])(regeneratorRuntime.mark((function e(t){var n,a;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(n=localStorage.getItem("sessionKey"),a=localStorage.getItem("userSecret"),!n||!a||""===a||""===n){e.next=6;break}return console.log("Keys found in storage"),this.commit("setLoginState",!0),e.abrupt("return",!0);case 6:return console.log("Keys not found in storage"),e.abrupt("return",!1);case 8:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}()}]),n}($["a"]),le=new $["c"]({namespaced:!1,state:oe,getters:ie,mutations:ce,actions:ue}),he=(n("4160"),n("a434"),n("159b"),function e(){Object(G["a"])(this,e),this.id="",this.username="",this.email="",this.pool="",this.token="",this.userSecret="",this.allUsers=[]}),de=function(e){Object(W["a"])(n,e);var t=Object(H["a"])(n);function n(){return Object(G["a"])(this,n),t.apply(this,arguments)}return Object(J["a"])(n,[{key:"username",value:function(){return this.state.username}},{key:"email",value:function(){return this.state.email}},{key:"pool",value:function(){return this.state.pool}},{key:"allUsers",value:function(){return this.state.allUsers}},{key:"getNewToken",value:function(e){return se.generateToken({username:this.state.username,timeStamp:e,sessionKey:this.state.token,userSecret:this.state.userSecret})}}]),n}($["b"]),pe=function(e){Object(W["a"])(n,e);var t=Object(H["a"])(n);function n(){return Object(G["a"])(this,n),t.apply(this,arguments)}return Object(J["a"])(n,[{key:"setToken",value:function(e){null!==e?this.state.token=e:console.log("Error!!! Token not found!")}},{key:"setUserSecret",value:function(e){null!==e?this.state.userSecret=e:console.log("Error!!! Token not found!")}},{key:"setUsername",value:function(e){null!==e&&(this.state.username=e)}},{key:"setEmail",value:function(e){this.state.email=e}},{key:"setPool",value:function(e){this.state.pool=e}},{key:"setAllUsers",value:function(e){var t=this;this.state.allUsers.splice(0,this.state.allUsers.length),e.forEach((function(e){t.state.allUsers.push(e)}))}}]),n}($["d"]),fe=function(e){Object(W["a"])(n,e);var t=Object(H["a"])(n);function n(){return Object(G["a"])(this,n),t.apply(this,arguments)}return Object(J["a"])(n,[{key:"loadDataLocalStorage",value:function(){this.commit("setUsername",localStorage.getItem("username")),this.commit("setToken",localStorage.getItem("sessionKey")),this.commit("setUserSecret",localStorage.getItem("userSecret"))}},{key:"initUserState",value:function(){var e=Object(q["a"])(regeneratorRuntime.mark((function e(){var t,n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return this.dispatch("loadDataLocalStorage"),t=Date.now().toString(),e.next=4,te.post("user",{data:null,auth:{username:this.state.username,token:se.generateToken({username:this.state.username,timeStamp:t,sessionKey:this.state.token,userSecret:this.state.userSecret})}},t);case 4:if(n=e.sent,!0===n.status){e.next=8;break}return console.log("Failed to get user data!"),e.abrupt("return");case 8:return this.commit("setEmail",n.data.email),this.commit("setPool",n.data.poolId),this.dispatch("loadAllUsers"),e.next=13,Te.dispatch("setAllChats",n.data.chats);case 13:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}()},{key:"loadAllUsers",value:function(){var e=Object(q["a"])(regeneratorRuntime.mark((function e(){var t,n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return this.dispatch("loadDataLocalStorage"),t=Date.now().toString(),e.next=4,te.post("all_users",{data:null,auth:{username:this.state.username,token:se.generateToken({username:this.state.username,timeStamp:t,sessionKey:this.state.token,userSecret:this.state.userSecret})}},t);case 4:if(n=e.sent,!0===n.status){e.next=8;break}return console.log("Failed to get users!"),e.abrupt("return");case 8:this.commit("setAllUsers",n.data);case 9:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}()}]),n}($["a"]),ge=new $["c"]({namespaced:!1,state:he,getters:de,mutations:pe,actions:fe}),me=(n("7db0"),n("c740"),n("d81d"),n("b0c0"),n("3ca3"),n("ddb0"),n("6821")),ve=n.n(me),be=function e(){Object(G["a"])(this,e),this.isNew=!1,this.chats=[],this.currentChat={id:"",admin:"",name:"",usernames:[],messages:[]},this.generateMessageId=function(){return ve()(Date.now()+Math.random())}},we=function(e){Object(W["a"])(n,e);var t=Object(H["a"])(n);function n(){return Object(G["a"])(this,n),t.apply(this,arguments)}return Object(J["a"])(n,[{key:"getCurrentChat",value:function(){var e=this;return this.state.isNew?this.state.currentChat:this.state.chats.find((function(t){return t.id===e.state.currentChat.id}))}},{key:"getAllChats",value:function(){return this.state.chats}},{key:"isNew",value:function(){return this.state.isNew}}]),n}($["b"]),ye=function(e){Object(W["a"])(n,e);var t=Object(H["a"])(n);function n(){return Object(G["a"])(this,n),t.apply(this,arguments)}return Object(J["a"])(n,[{key:"setChats",value:function(e){var t=this;this.state.chats.splice(0,this.state.chats.length),e.forEach((function(e){t.state.chats.push(e)}))}},{key:"addMessageToItsChat",value:function(e){this.state.chats.forEach((function(t){if(t.id===e.chatId){var n=t.messages.findIndex((function(t){return t.id===e.id}));n<0?t.messages.push(e):(t.messages[n].text=e.text,t.messages[n].state=e.state)}}))}},{key:"addChat",value:function(e){console.log("Adding chat: ",e),void 0===e.messages&&(e.messages=new Array),void 0===this.state.chats.find((function(t){return t.id===e.id}))&&this.state.chats.push(e)}},{key:"setCurrentChat",value:function(e){this.state.currentChat=e}},{key:"setNewValue",value:function(e){this.state.isNew=e}}]),n}($["d"]),ke=function(e){Object(W["a"])(n,e);var t=Object(H["a"])(n);function n(){return Object(G["a"])(this,n),t.apply(this,arguments)}return Object(J["a"])(n,[{key:"sendMessageInChat",value:function(){var e=Object(q["a"])(regeneratorRuntime.mark((function e(t){var n,a,r,s;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(!this.state.isNew){e.next=12;break}return n=Date.now().toString(),e.next=4,te.post("chat",{data:this.state.currentChat,auth:{username:Te.getters.username(),token:Te.getters.getNewToken(n)}},n);case 4:if(a=e.sent,!0===a.status){e.next=8;break}return console.log("Failed to send init message! Failed to create chat!"),e.abrupt("return");case 8:r={id:a.data.id,admin:a.data.admin,name:a.data.name,usernames:a.data.usernames,messages:new Array},this.commit("addChat",r),this.commit("setCurrentChat",r),this.commit("setNewValue",!1);case 12:return s={id:this.state.generateMessageId(),sender:Te.getters.username(),chatId:this.state.currentChat.id,date:Date.now(),state:1,text:t,attachedFilePath:""},this.commit("addMessageToItsChat",s),e.next=16,Te.dispatch("sendSocketMessage",{type:1,message:s});case 16:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}()},{key:"setCurrentChat",value:function(){var e=Object(q["a"])(regeneratorRuntime.mark((function e(t){var n,a,r;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:t.isNew?(n=t.data,a={id:"",name:Te.getters.username()+" and "+n.username,usernames:[Te.getters.username(),n.username],admin:Te.getters.username(),messages:[]},this.commit("setNewValue",!0),this.commit("setCurrentChat",a)):(r=t.data,this.commit("setCurrentChat",r),this.commit("setNewValue",!1),r.messages.forEach((function(e){e.sender!==Te.getters.username()&&3!==e.state&&(e.state=3,Te.dispatch("sendSocketMessage",{type:2,message:e}))})));case 1:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}()},{key:"setAllChats",value:function(){var e=Object(q["a"])(regeneratorRuntime.mark((function e(t){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,Promise.all(t.map(function(){var e=Object(q["a"])(regeneratorRuntime.mark((function e(t){var n,a;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return n=Date.now().toString(),e.next=3,te.post("messages",{data:t,auth:{username:Te.getters.username(),token:Te.getters.getNewToken(n)}},n);case 3:if(a=e.sent,!0===a.status){e.next=7;break}return console.log("Error loading messages"),e.abrupt("return");case 7:t.messages=a.data?a.data:[];case 8:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}()));case 2:this.commit("setChats",t);case 3:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}()}]),n}($["a"]),je=new $["c"]({namespaced:!1,state:be,getters:we,mutations:ye,actions:ke}),Oe=n("7e3c").w3cwebsocket,Se=1,Ce=function e(){Object(G["a"])(this,e),this.isConnected=!1},xe=function(e){Object(W["a"])(n,e);var t=Object(H["a"])(n);function n(){return Object(G["a"])(this,n),t.apply(this,arguments)}return Object(J["a"])(n,[{key:"isConnected",value:function(){return this.state.isConnected}}]),n}($["b"]),Le=function(e){Object(W["a"])(n,e);var t=Object(H["a"])(n);function n(){return Object(G["a"])(this,n),t.apply(this,arguments)}return Object(J["a"])(n,[{key:"setClient",value:function(e){this.state.client=e}},{key:"setConnected",value:function(e){this.state.isConnected=e}}]),n}($["d"]),Ie=function(e){Object(W["a"])(n,e);var t=Object(H["a"])(n);function n(){return Object(G["a"])(this,n),t.apply(this,arguments)}return Object(J["a"])(n,[{key:"sendSocketMessage",value:function(e){void 0!==this.state.client?this.state.client.send(JSON.stringify(e)):console.log("Error! Client undefined!")}},{key:"initWebSocket",value:function(){var e=this;if(!this.state.isConnected){var t=Date.now().toString(),n=Te.getters.username(),a=Te.getters.getNewToken(t),r=new Oe("ws://localhost:8080/api/v1/connect?user="+n+"&token="+a+"&time="+t,"chat");r.onopen=function(){e.dispatch("onOpen")},r.onclose=function(){e.dispatch("onClose")},r.onmessage=function(t){e.dispatch("onMessage",t.data)},r.onerror=function(){console.log("Connection Error")},this.commit("setClient",r)}}},{key:"onOpen",value:function(){console.log("Opened ws"),this.commit("setConnected",!0)}},{key:"onClose",value:function(){console.log("Closed ws"),this.commit("setConnected",!1)}},{key:"onMessage",value:function(e){console.log("Got message in ws: ",e);var t=JSON.parse(e);if(t.type===Se)Te.commit("addMessageToItsChat",t.data);else{var n=Date.now().toString();te.post("messages",{data:t.data,auth:{username:Te.getters.username(),token:Te.getters.getNewToken(n)}},n).then((function(e){if(!0===e.status){var n=t.data,a={id:n.id,admin:n.admin,name:n.name,usernames:n.usernames,messages:new Array};Array.isArray(e.data)&&(a.messages=e.data),Te.commit("addChat",t.data)}else console.log("Error loading messages")}))}}}]),n}($["a"]),_e=new $["c"]({namespaced:!1,state:Ce,getters:xe,mutations:Le,actions:Ie});a["a"].use(B);var Ae=new $["c"]({modules:{Login:le,UserData:ge,Chats:je,WebSocket:_e}}),Te=Object($["e"])(Ae,{strict:!1}),Re=(n("d1e78"),n("f309")),Ee=n("0879");a["a"].use(Re["a"]);var Ne=new Re["a"]({theme:{themes:{light:{primary:"#ee44aa",secondary:"#424242",accent:"#82B1FF",error:"#FF5252",info:"#2196F3",success:"#4CAF50",warning:"#FFC107"}}},lang:{locales:{ru:Ee["a"]},current:"ru"}});a["a"].config.productionTip=!1,new a["a"]({router:K,store:Te,vuetify:Ne,render:function(e){return e(p)}}).$mount("#app")},e118:function(e,t,n){},eeb1:function(e,t,n){"use strict";var a=n("e118"),r=n.n(a);r.a}});
//# sourceMappingURL=app.de1d4eed.js.map