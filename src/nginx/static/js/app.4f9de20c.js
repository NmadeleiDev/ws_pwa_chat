(function(t){function e(e){for(var n,o,i=e[0],c=e[1],u=e[2],d=0,f=[];d<i.length;d++)o=i[d],Object.prototype.hasOwnProperty.call(a,o)&&a[o]&&f.push(a[o][0]),a[o]=0;for(n in c)Object.prototype.hasOwnProperty.call(c,n)&&(t[n]=c[n]);l&&l(e);while(f.length)f.shift()();return r.push.apply(r,u||[]),s()}function s(){for(var t,e=0;e<r.length;e++){for(var s=r[e],n=!0,i=1;i<s.length;i++){var c=s[i];0!==a[c]&&(n=!1)}n&&(r.splice(e--,1),t=o(o.s=s[0]))}return t}var n={},a={app:0},r=[];function o(e){if(n[e])return n[e].exports;var s=n[e]={i:e,l:!1,exports:{}};return t[e].call(s.exports,s,s.exports,o),s.l=!0,s.exports}o.m=t,o.c=n,o.d=function(t,e,s){o.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:s})},o.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},o.t=function(t,e){if(1&e&&(t=o(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var s=Object.create(null);if(o.r(s),Object.defineProperty(s,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var n in t)o.d(s,n,function(e){return t[e]}.bind(null,n));return s},o.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return o.d(e,"a",e),e},o.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},o.p="/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],c=i.push.bind(i);i.push=e,i=i.slice();for(var u=0;u<i.length;u++)e(i[u]);var l=c;r.push([0,"chunk-vendors"]),s()})({0:function(t,e,s){t.exports=s("56d7")},"034f":function(t,e,s){"use strict";var n=s("85ec"),a=s.n(n);a.a},"0a35":function(t,e,s){"use strict";var n=s("c1a1"),a=s.n(n);a.a},"56d7":function(t,e,s){"use strict";s.r(e);s("e260"),s("e6cf"),s("cca6"),s("a79d");var n=s("2b0e"),a=s("2f62"),r=(s("96cf"),s("1da1")),o=s("bc3a"),i=s.n(o),c=function(){var t=Object(r["a"])(regeneratorRuntime.mark((function t(e){var s,n,a;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return s=l(e),t.next=3,i.a.get(s);case 3:return a=t.sent,a.data.data,n=a.data.status,console.log("GET: resultStatus: ",n),!0!==n&&console.log(a),t.abrupt("return",a.data);case 9:case"end":return t.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}(),u=function(){var t=Object(r["a"])(regeneratorRuntime.mark((function t(e){var s,n,a,r=arguments;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return s=r.length>1&&void 0!==r[1]?r[1]:null,n=l(e),t.next=4,i.a.post(n,s);case 4:return a=t.sent,console.log("POST: resultStatus: ",a.data.status),!0!==a.data.status&&console.log("failed response: ",a),t.abrupt("return",a.data);case 8:case"end":return t.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}();function l(t){return"/api/v1/"+t}var d={get:c,post:u};n["default"].use(a["a"]);var f=new a["a"].Store({state:{user:{username:"",email:"",contacts:[],chats:[{chat_id:"",name:"",usernames:"",admin:""}]},allUsers:[],socket:{isConnected:!1,message:"",reconnectError:!1}},getters:{GET_USER:function(t){return t.user},GET_ALL_USERS:function(t){return t.allUsers}},mutations:{SET_USER:function(t,e){t.user=e,console.log("User set: ",t.user)},SET_ALL_USERS:function(t,e){t.allUsers=e,console.log("All users set: ",t.allUsers)},SOCKET_ONOPEN:function(t,e){n["default"].prototype.$socket=e.currentTarget,t.socket.isConnected=!0},SOCKET_ONCLOSE:function(t,e){t.socket.isConnected=!1},SOCKET_ONERROR:function(t,e){console.error(t,e)},SOCKET_ONMESSAGE:function(t,e){console.log("Got message in store: ",e),t.socket.message=e},SOCKET_RECONNECT:function(t,e){console.info(t,e)},SOCKET_RECONNECT_ERROR:function(t){t.socket.reconnectError=!0}},actions:{LOAD_USER_DATA:function(t){var e=d.get("get_data");e.then((function(e){t.commit("SET_USER",e.data)}))},LOAD_ALL_USERS:function(t){var e=d.get("all_users");e.then((function(e){t.commit("SET_ALL_USERS",e.data)}))},SEND_MESSAGE:function(t,e){n["default"].prototype.$socket.send(e)},RECEIVE_MESSAGE:function(t,e){t.commit("SOCKET_ONMESSAGE",e)}}}),p=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{attrs:{id:"app"}},[s("router-view")],1)},h=[],m={name:"App",components:{}},g=m,b=(s("034f"),s("2877")),v=Object(b["a"])(g,p,h,!1,null,null,null),_=v.exports,w=(s("f9e3"),s("2dd8"),s("8c4f")),E=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("b-container",{attrs:{fluid:""}},[s("b-alert",{attrs:{variant:"danger",dismissible:""},model:{value:t.showAlert,callback:function(e){t.showAlert=e},expression:"showAlert"}},[t._v(" "+t._s(t.alertText)+" ")]),s("b-row",{staticClass:"m-4 mt-5"},[s("b-col",[s("h1",[t._v("Welcome back")])])],1),s("b-row",{staticClass:"m-4"},[s("b-col",[s("label",[t._v(" Имя "),s("b-form-input",{attrs:{size:"lg",placeholder:"Enter your nick"},model:{value:t.user.username,callback:function(e){t.$set(t.user,"username",e)},expression:"user.username"}})],1)])],1),s("b-row",{staticClass:"m-4"},[s("b-col",[s("label",[t._v(" Пароль "),s("b-form-input",{attrs:{size:"lg",state:t.passwordOk,type:"password",placeholder:"Enter your password"},model:{value:t.user.password,callback:function(e){t.$set(t.user,"password",e)},expression:"user.password"}})],1)])],1),s("b-collapse",{staticClass:"mt-2",attrs:{id:"collapse-1"}},[s("b-row",{staticClass:"m-4"},[s("b-col",[s("label",[t._v(" Пароль еще раз "),s("b-form-input",{attrs:{size:"lg",state:t.passwordOk,type:"password",placeholder:"Enter password again"},model:{value:t.passwordVerif,callback:function(e){t.passwordVerif=e},expression:"passwordVerif"}})],1)])],1)],1),s("b-row",{staticClass:"m-4 mt-5"},[s("b-col",[s("b-button",{attrs:{variant:"success",size:"lg",block:t.isMobile},on:{click:function(e){t.isCreate?t.signUp():t.signIn()}}},[t._v(" "+t._s(t.isCreate?"Создать аккаут":"Войти")+" ")])],1)],1),s("b-row",[s("b-col",[s("b-button",{directives:[{name:"b-toggle",rawName:"v-b-toggle.collapse-1",modifiers:{"collapse-1":!0}}],attrs:{variant:"outline-secondary"},on:{click:t.changeAction}},[t._v(" "+t._s(t.isCreate?"У меня есть аккаут":"У меня нет аккаутна")+" ")])],1)],1)],1)},S=[],C={name:"Login",data:function(){return{user:{username:"",password:""},passwordVerif:"",isCreate:!1,showAlert:!1,alertText:!1,isMobile:!0}},created:function(){this.isMobile=window.screen.width>500},methods:{signIn:function(){var t=this,e=this,s=d.post("signin",this.user);s.then((function(s){!0===s.status?e.$router.push("/chat"):(t.showAlert=!0,t.alertText="Имя и пароль не совпадают.")}))},signUp:function(){var t=this,e=this,s=d.post("signup",this.user);s.then((function(s){!0===s.status?e.$router.push("/chat"):(t.showAlert=!0,t.alertText="Не удалось создать аккаунт. Попробуйте позже.")}))},changeAction:function(){this.isCreate=!this.isCreate}},computed:{passwordOk:function(){return this.isCreate&&0!==this.user.password.length?this.user.password===this.passwordVerif:null}}},x=C,O=Object(b["a"])(x,E,S,!1,null,"247186ab",null),T=O.exports,k=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("b-container",[s("b-sidebar",{attrs:{id:"sidebar-1",title:"Chats",shadow:"",backdrop:""},scopedSlots:t._u([{key:"footer",fn:function(e){var n=e.hide;return[s("div",{staticClass:"d-flex bg-dark text-light justify-content-around align-items-center px-3 py-2"},[s("b-button",{attrs:{variant:"outline-warning"},on:{click:t.signOut}},[t._v("Выйти")]),s("b-button",{attrs:{size:"md"},on:{click:n}},[t._v("Закрыть")])],1)]}}])},[s("b-container",[s("b-row",{staticClass:"m-3 mt-5"},[s("b-col",{attrs:{cols:"4"}},[s("b-avatar",{attrs:{variant:"primary",text:"BV"}})],1),s("b-col",{attrs:{cols:"6"}},[s("h3",[t._v(t._s(t.userData.username))])])],1),s("b-row",[s("b-col",{staticClass:"m-3"},[s("b-button",{directives:[{name:"b-toggle",rawName:"v-b-toggle.all-chats-cont",modifiers:{"all-chats-cont":!0}}],staticClass:"mb-1",attrs:{variant:"info",block:""}},[t._v("Активные чаты")]),s("b-collapse",{attrs:{id:"all-chats-cont"}},[t.userData.chats.length>0?s("b-list-group",t._l(t.userData.chats,(function(e){return s("b-list-group-item",{key:e.chat_id,on:{click:function(s){return t.setExistingChat(e)}}},[t._v(" "+t._s(e.name)+" ")])})),1):s("p",[t._v("Чатов пока нет...")])],1)],1)],1),s("b-row",[s("b-col",{staticClass:"m-3"},[s("b-button",{directives:[{name:"b-toggle",rawName:"v-b-toggle.all-users-cont",modifiers:{"all-users-cont":!0}}],staticClass:"mb-1",attrs:{variant:"info",block:""}},[t._v("Найти пользователя")]),s("b-collapse",{attrs:{id:"all-users-cont"}},[s("b-input",{attrs:{placeholder:"Поиск пользователей..."},model:{value:t.findUserText,callback:function(e){t.findUserText=e},expression:"findUserText"}}),t.newUsers.length>0?s("b-list-group",t._l(t.newUsers.filter((function(e){return e.username.includes(t.findUserText)>0})),(function(e){return s("b-list-group-item",{key:e.username,staticClass:"flex-row align-items-start",on:{click:function(s){return t.setNewChat(e)}}},[s("div",{staticClass:"d-flex flex-row align-content-center justify-content-around"},[s("div",{staticClass:"align-self-lg-start"},[s("b-avatar")],1),s("div",[s("h4",[t._v(t._s(e.username))])])])])})),1):t._e()],1)],1)],1)],1)],1),s("b-button",{directives:[{name:"b-toggle",rawName:"v-b-toggle.sidebar-1",modifiers:{"sidebar-1":!0}}],staticClass:"position-absolute fixed-top",attrs:{squared:"",variant:"primary"}},[t._v("Меню")]),s("b-row",[s("b-col",[s("h4",[t._v("Чат с "+t._s(t.chat.usernames.join(", ")))])])],1),s("b-row",[s("b-col",t._l(t.messages,(function(e){return s("div",{key:e.date,class:(e.sender===t.userData.username?"right-message":"left-message")+" message-container"},[s("div",{staticClass:"message-header"},[s("div",{staticClass:"message-sender"},[t._v(" "+t._s(e.sender)+" ")]),s("div",{staticClass:"message-date"},[t._v(" "+t._s(t.convertToDatetime(e.date))+" ")])]),s("div",{staticClass:"message-text"},[t._v(t._s(e.text))])])})),0)],1),s("b-row",{attrs:{"align-content":"center"}},[s("b-col",{attrs:{"align-self":"center"}},[s("div",{staticClass:"input-container w-75 m-auto"},[s("b-input-group",[s("b-input",{staticClass:"message-input",attrs:{placeholder:"Введите сообщение...",type:"text"},model:{value:t.messageText,callback:function(e){t.messageText=e},expression:"messageText"}}),s("b-input-group-append",[s("b-button",{attrs:{variant:"outline-primary"},on:{click:t.sendMessage}},[s("b-icon-envelope")],1)],1)],1)],1)])],1)],1)},y=[],A=(s("4de4"),s("4160"),s("caad"),s("b0c0"),s("2532"),s("159b"),{name:"Chat",data:function(){return{chat:{chat_id:null,name:"",usernames:[],admin:""},messages:[],messageText:"",findUserText:"",showSidebar:!1}},created:function(){this.$store.dispatch("LOAD_USER_DATA"),this.$store.dispatch("LOAD_ALL_USERS"),this.connection=this.setConnection()},methods:{sendMessage:function(){var t={sender:this.userData.username,chat_id:this.chat.chat_id,meta:this.chat.name,date:Date.now(),state:0,text:this.messageText,attached_file_path:null};console.log("Message to send: ",t),this.messageText="",this.$store.dispatch("SEND_MESSAGE",JSON.stringify(t))},setConnection:function(){var t=this;this.$connect(),this.$options.sockets.onmessage=function(e){var s=JSON.parse(e.data);t.$store.dispatch("RECEIVE_MESSAGE",s),t.messages.push(s)}},setExistingChat:function(t){this.chat=t;var e=d.get("get_messages/"+t.chat_id),s=this;this.messages=[],e.then((function(t){void 0!==t&&null!==t&&!1!==t.status&&0!==t.data.length&&t.data.forEach((function(t){s.messages.push(t)}))}))},setNewChat:function(t){this.chat={chat_id:null,name:t.username,usernames:[t.username,this.userData.username],admin:this.userData.username},this.messages=[]},convertToDatetime:function(t){var e=new Date(t),s=e.getHours(),n="0"+e.getMinutes(),a="0"+e.getSeconds();return s+":"+n.substr(-2)+":"+a.substr(-2)},signOut:function(){d.get("signout"),this.$router.push("/login")}},computed:{userData:function(){return this.$store.getters.GET_USER},allUsers:function(){return this.$store.getters.GET_ALL_USERS},newUsers:function(){var t=this;return this.$store.getters.GET_ALL_USERS.length<1?[]:this.$store.getters.GET_ALL_USERS.filter((function(e){return e.username!==t.userData.username&&(null==t.userData.contacts||!t.userData.contacts.includes(e.username))}))}}}),U=A,R=(s("0a35"),Object(b["a"])(U,k,y,!1,null,"b3a416b6",null)),D=R.exports;n["default"].use(w["a"]);var L=[{path:"/",redirect:"/login"},{path:"/login",name:"login_page",meta:{title:"login page"},component:T},{path:"/chat",name:"chat_page",meta:{title:"Chat page"},component:D}],$=new w["a"]({mode:"history",base:"/",routes:L}),N=$,M=s("5f5b"),j=s("b1e0"),G=s("b408"),P=s.n(G);n["default"].use(P.a,"ws://"+window.location.host+"/ws/connect",{connectManually:!0,reconnection:!0,reconnectionAttempts:5},{store:f}),n["default"].use(M["a"]),n["default"].use(j["a"]),n["default"].config.productionTip=!1,new n["default"]({store:f,router:N,render:function(t){return t(_)}}).$mount("#app")},"85ec":function(t,e,s){},c1a1:function(t,e,s){}});
//# sourceMappingURL=app.4f9de20c.js.map