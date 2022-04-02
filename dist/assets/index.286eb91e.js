var J=Object.defineProperty;var j=Object.getOwnPropertySymbols;var K=Object.prototype.hasOwnProperty,Y=Object.prototype.propertyIsEnumerable;var D=(s,t,n)=>t in s?J(s,t,{enumerable:!0,configurable:!0,writable:!0,value:n}):s[t]=n,R=(s,t)=>{for(var n in t||(t={}))K.call(t,n)&&D(s,n,t[n]);if(j)for(var n of j(t))Y.call(t,n)&&D(s,n,t[n]);return s};import{a as z,u as M,r as h,j as o,G as B,B as w,P as k,C as W,b as L,c as e,T as C,d as q,e as E,f as S,D as I,S as T,F as _,I as Q,g as X,M as Z,h as b,i as ee,k as te,l as ne,m as se,n as re,L as V,o as U,p as ae,q as ie,s as $,t as oe,v as le,w as ce,x as ue,y as de,z as he,A as O,E as me,H as fe,J as pe,R as ge,K as be,N as ve,O as P}from"./vendor.b8856755.js";const xe=function(){const t=document.createElement("link").relList;if(t&&t.supports&&t.supports("modulepreload"))return;for(const r of document.querySelectorAll('link[rel="modulepreload"]'))i(r);new MutationObserver(r=>{for(const a of r)if(a.type==="childList")for(const u of a.addedNodes)u.tagName==="LINK"&&u.rel==="modulepreload"&&i(u)}).observe(document,{childList:!0,subtree:!0});function n(r){const a={};return r.integrity&&(a.integrity=r.integrity),r.referrerpolicy&&(a.referrerPolicy=r.referrerpolicy),r.crossorigin==="use-credentials"?a.credentials="include":r.crossorigin==="anonymous"?a.credentials="omit":a.credentials="same-origin",a}function i(r){if(r.ep)return;r.ep=!0;const a=n(r);fetch(r.href,a)}};xe();var N=z.create({baseURL:"",headers:{"Content-Type":"application/json"}}),x={async request(s,t,n={}){const i={"Content-Type":"application/json"};try{let r;switch(s){case"get":r=await N.get(encodeURI(t),{headers:i,params:n||{}});break;case"patch":r=await N.patch(encodeURI(t),n,{headers:i});break;case"post":r=await N.post(encodeURI(t),n,{headers:i});break;case"put":r=await N.put(encodeURI(t),n,{headers:i});break;case"delete":r=await N.delete(encodeURI(t),{headers:i});break;default:throw new Error("APIRequests: Unsupported Request Type")}return r}catch(r){return r.response}}};const y="/s";var A={async getHealth(){return await x.request("get","")},getPort(){return x.request("get","port")},async getSession(s){return await x.request("get",`${y}/${s}`)},async getSessions(){return await x.request("get",y)},async postSession(s){return await x.request("post",y,s)},async postPoints(s,t){return await x.request("post",`${y}/${s}`,t)},async patchSession(s,t){return await x.request("patch",`${y}/${s}`,t)},async deleteAction(s){return await x.request("delete",`${y}/${s}`)}};function Ce(){const s=M(),[t,n]=h.exports.useState(""),[i,r]=h.exports.useState(!1),[a,u]=h.exports.useState(!1);h.exports.useEffect(()=>{A.getHealth().then(l=>{l||u(!0)})});async function d(){if(!t)return null;await A.getSession(Number(t)).then(l=>{l&&l.status===200?s(`/room/${t}`):(n(""),r(!0))})}return o(B,{container:!0,spacing:0,direction:"column",alignItems:"center",justifyContent:"center",style:{minHeight:"100vh"},children:[o(w,{children:[o(k,{sx:{minWidth:1e3,minHeight:600},elevation:0,children:[o(W,{elevation:6,children:[o(L,{children:[e(C,{variant:"h5",children:"Go to the room \u{1F575}\uFE0F"}),e("br",{}),e(q,{label:"room id",variant:"standard",value:t,onChange:l=>{n(l.target.value)},fullWidth:!0})]}),e(E,{children:e(S,{variant:"contained",onClick:()=>{d()},disabled:!t,children:"GO!"})})]}),e("br",{}),e("br",{}),e(I,{children:"\u{1F446} or \u{1F447}"}),e("br",{}),e("br",{}),e(w,{textAlign:"center",children:e(S,{onClick:()=>s("/room"),children:"Create new one \u{1F929}"})})]}),")"]}),e(T,{sx:{position:"fixed",bottom:0,left:0,right:0},open:i,autoHideDuration:5e3,onClose:()=>r(!1),message:"Can not find a room \u{1F614} \u{1F6AB}"})]})}class Se extends h.exports.Component{constructor(t){super(t);this.state={}}async componentDidMount(){await A.getSessions().then(t=>{t&&this.setState({sessions:t.data})})}render(){return e("div",{className:"App",children:e(Ce,{})})}}function ye(){const[s,t]=h.exports.useState(""),[n,i]=h.exports.useState("10"),[r,a]=h.exports.useState(!1),u=M(),d=6e3,l=()=>{a(!0)},c=()=>{a(!1)};async function m(){let g={name:s};await A.postSession(g).then(v=>{v.status===404?l():v.status===201&&u(`/room/${v.data.id}`)}),c()}return e(B,{container:!0,spacing:0,direction:"column",alignItems:"center",justifyContent:"center",style:{minHeight:"100vh"},children:o(w,{children:[e(k,{elevation:5,children:o(W,{sx:{maxWidth:600,mx:"auto"},children:[o(L,{children:[e(C,{gutterBottom:!0,variant:"h5",children:"Create a new room"}),e(C,{variant:"body2",color:"text.secondary",children:"Provide a name for your room and the voting system."}),e(I,{}),e("div",{children:e(_,{fullWidth:!0,children:e(q,{sx:{maxWidth:600},id:"text-field-room-name",label:"Room name",variant:"standard",required:!0,fullWidth:!0,autoFocus:!0,onChange:g=>{t(g.target.value)}})})}),e("br",{}),e("div",{children:o(_,{fullWidth:!0,children:[e(Q,{id:"select-input-label",children:"Voting system"}),e(X,{labelId:"select-input-label",id:"select-input-voting",value:n,label:"Voting system",onChange:g=>i(g.target.value),variant:"standard",children:e(Z,{value:10,children:"Fibonacii (1, 2, 3, 5, 8, 13)"})})]})})]}),o(E,{children:[e(S,{size:"large",color:"primary",disabled:!s,onClick:m,children:"Create"}),e(S,{color:"secondary",onClick:()=>u("/"),children:"Cancel"})]})]})}),e(T,{open:r,autoHideDuration:d,onClose:()=>a(!1),message:"Can not create a session"})]})})}function we(s){const{session_id:t,nickname:n,gws:i}=s,r=[0,1,2,3,5,8,13,21],a=d=>{s.gws.send(JSON.stringify({action:"salle",payload:[n,d].join(",")}))};return e(b,{children:e(()=>o(b,{children:[e("br",{}),e(ee,{direction:"row",spacing:0,divider:e(I,{orientation:"vertical",flexItem:!0}),alignItems:"center",justifyContent:"center",children:r.map((d,l)=>e(S,{sx:{maxWidth:120},onClick:()=>a(d),children:d},l))})]}),{})})}function ke(s){const{onClose:t,open:n}=s,[i,r]=h.exports.useState(""),a=(d,l)=>{l&&l=="backdropClick"||t(i)},u=d=>{t(d)};return o(te,{onClose:a,open:n,sx:{maxWidth:"md",mx:"auto"},fullWidth:!0,disableEscapeKeyDown:!0,children:[e(ne,{children:"Who are you?"}),e(se,{children:e(q,{id:"standard-basic",label:"Your nickname \u{1F60E}",variant:"standard",onChange:d=>r(d.target.value.trim()),fullWidth:!0,autoFocus:!0})}),e(re,{children:e(S,{onClick:()=>u(i),children:"GO!"})})]})}function Ie(s){const{users:t,minimum:n,maximum:i}=s;let r=0,a=0;return t.forEach(u=>{(u.salle?u.salle:-1)>=0&&(r+=u.salle,a+=1)}),e(b,{children:a>0&&e(b,{children:n!==i&&o(b,{children:["\u{1F916} Avarage: ",e("b",{children:(r/a).toFixed(2)})]})})})}function Fe(s){if(!s.users)return null;const t=s.users,n=a(t),i=Math.min(...n),r=Math.max(...n);function a(l){return l.filter(c=>c).map(c=>c.salle?c.salle:-1).filter(c=>c>0)}function u(l,c){return l.filter(m=>m).map(m=>(m.salle?m.salle:-1)===c?m.username:"").filter(m=>m)}const d=()=>{if(!t)return null;let l=0,c=0,m,g;const v=a(t);return l=Math.min(...v),c=Math.max(...v),m=u(t,l),g=u(t,c),o(b,{children:[l!==c&&o(b,{children:["\u{1F4C8} max: ",e("b",{children:c}),e("br",{}),"\u2003(",g.join(","),")",e("br",{}),"\u{1F4C9} min:\xA0 ",e("b",{children:l}),e("br",{}),"\u2003(",m.join(","),")",e("br",{}),e("br",{}),e(w,{children:e(ae,{getAriaLabel:()=>"Points range",value:[l,c],valueLabelDisplay:"auto",marks:t.map(f=>({value:f.salle?Number(f.salle):0,label:(f.salle?Number(f.salle):0).toString()})).filter(f=>f.value>0),disabled:!0,min:0,max:21})})]}),l===c&&o(b,{children:[" ","\u{1F44F} \u{1F64C} full consent for ",e("b",{children:l}),"!"]})]})};return o(k,{elevation:0,children:[a(t).length>0&&s.showValues&&o(V,{subheader:e(U,{children:"Statisitcs"}),children:[e(Ie,{users:t,minimum:i,maximum:r}),e("br",{}),e(d,{})]}),a(t).length<=0&&e(b,{children:'Waiting for "your" score! \u{1F44A}'})]})}function Ne(s){const{reveal:t,session_id:n,ws:i}=s;async function r(u){await A.patchSession(Number(n),{reveal:u}),i.send(JSON.stringify({action:"reveal",payload:u}))}async function a(){await r(!1),i.send(JSON.stringify({action:"restart",payload:null}))}return e(k,{sx:{position:"fixed",bottom:0,left:0,right:0},elevation:3,children:o(ie,{showLabels:!0,children:[e($,{onClick:()=>{i.send("")},label:"Refresh",icon:e(oe,{})}),!t&&e($,{onClick:()=>{r(!t)},label:"Show values",icon:e(le,{})}),t&&e($,{onClick:()=>{r(!t)},label:"Hide values",icon:e(ce,{})}),e($,{onClick:()=>{a()},label:"Replay",icon:e(ue,{})})]})})}function G(s){let t=0,n;for(n=0;n<s.length;n+=1)t=s.charCodeAt(n)+((t<<5)-t);let i="#";for(n=0;n<3;n+=1)i+=`00${(t>>n*8&255).toString(16)}`.slice(-2);return i}function Ae(s,t){const n=s.split(" "),i=n[0][0],r=()=>n.length>1?n[1][0]:"";return{sx:{bgcolor:G(t||"")},children:`${i} ${r()}`}}function H(s,t){return{sx:{bgcolor:G(t||"")},children:`${s}`}}function Re(s){function t(n,i){return n!==i}return o(w,{children:[e(I,{}),e(V,{sx:{width:"100%",bgcolor:"background.paper"},subheader:e(U,{component:"div",id:"nested-list-subheader",children:"Participants"}),children:s.session.users&&s.session.users.map(n=>{if((n.salle?n.salle:-1)>=0)return o(de,{alignItems:"flex-start",children:[o(he,{children:[s.showValues&&e(O,R({},H(n.salle,n.username))),!s.showValues&&(t(n.username,s.nickname)?e(O,R({},Ae("#",n.username))):e(O,R({},H(n.salle,n.username))))]}),e(me,{primary:e(h.exports.Fragment,{children:e(C,{sx:{display:"inline"},variant:"h5",children:o(C,{variant:"body1",children:[n.username," ",n.username===s.nickname&&e(C,{variant:"body2",children:"(You)"})]})})})})]},n.id)})})]})}function $e(s){const{nickname:t,session_name:n,room_id:i}=s;return o(k,{sx:{position:"fixed",top:10,left:10,right:0},elevation:0,children:["Welcome back, ",e("b",{children:t})," @ Session ",e("b",{children:n})," (#",i,")",e(I,{})]})}function We(){let s=fe();const[t,n]=h.exports.useState(!1),[i,r]=h.exports.useState(""),[a,u]=h.exports.useState({id:0,name:"",reveal:!1,users:[]}),[d,l]=h.exports.useState(),[c,m]=h.exports.useState(new WebSocket("ws://localhost")),g=p=>{n(!1),r(p),sessionStorage.setItem("nickname",p)},v=()=>{c.onopen=()=>{c.send(i||"")},c.onmessage=p=>{const F=JSON.parse(p.data);u(F)}},f=()=>a.users?a.users.filter(p=>(p.salle?p.salle:0)>0).length:0;return h.exports.useEffect(()=>{const p=sessionStorage.getItem("nickname");return p?r(p):!p&&!i&&n(!0),fetch("/port").then(F=>F.json()).then(F=>l(F.port)),m(new WebSocket(`ws://localhost${d}/ws/${s.roomId}`)),v(),()=>{}},["nickname"]),o(b,{children:[e($e,{nickname:i,session_name:a.name,room_id:s.roomId}),o(B,{container:!0,spacing:0,direction:"column",alignItems:"center",justifyContent:"center",style:{minHeight:"100vh"},children:[o(w,{sx:{maxWidth:1e3},alignContent:"center",justifyContent:"center",children:[o(k,{sx:{mx:"auto"},elevation:0,children:[o(W,{elevation:0,children:[o(L,{children:[a.users&&e(Fe,{users:a.users,showValues:a.reveal}),!a.reveal&&o(C,{variant:"body2",children:[f()==1?`\u{1F4EB} ${f()} response is hidden, `:" ",f()>1?`\u{1F4EB} ${f()} responses are hidden, `:" ","waiting for reveal... \u231B"," "]})]}),e(E,{children:e(S,{onClick:()=>{c.send("")},children:"refresh"})})]}),e("br",{}),e(W,{elevation:0,children:e(L,{children:e(Re,{nickname:i,showValues:a.reveal,session:a})})})]}),e(I,{}),e(we,{session_id:s.roomId,nickname:i,gws:c})]}),e(ke,{open:t,onClose:g})]}),e(Ne,{session_id:s.roomId,reveal:a.reveal,ws:c})]})}const Le=document.getElementById("root"),Oe=pe(Le);Oe.render(e(ge.StrictMode,{children:e(be,{children:o(ve,{children:[e(P,{path:"/",element:e(Se,{})}),e(P,{path:"room",element:e(ye,{})}),e(P,{path:"room/:roomId",element:e(We,{})})]})})}));
