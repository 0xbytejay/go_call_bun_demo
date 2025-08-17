import  koffi from 'koffi';

console.log("Start Running!");


const PrintCallback = koffi.proto('bool PrintCallback(uint16_t)');
let cb = koffi.register(function(arg){
 console.log(`js callback called from golang, value: ${arg}`);
 return arg % 2 ==0;
}, koffi.pointer(PrintCallback));


globalThis.registerCallback("jsCallbackDemo",Number(koffi.address(cb)))

let count = 0;
setInterval(() => {
  console.log(`Count: ${count++}`);
}, 1000);



// var ffi = require('ffi-napi');
// globalThis.testCustomValue = "Hello";

// console.log("Start Running!");

// // const jsCallbackDemo = new JSCallback(
// //    (arg) => {
// //     console.log(`js callback called from golang, value: ${arg}.......`);

// //     return true;
// //   },
// //   {
// //     returns: FFIType.bool,
// //     args: [FFIType.i16],
// //     threadsafe: true,
// //   },
// // );



// // globalThis.registerCallback("jsCallbackDemo",jsCallbackDemo.ptr)

// let count = 0;
// setInterval(() => {
//   console.log(`Count: ${count++}`);
// }, 1000);
