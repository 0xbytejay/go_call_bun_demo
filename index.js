console.log("Start Running!");

globalThis.testCustomValue = "Hello";


// globalThis.testFunc = () => {
//   console.log("Call testFunc");
// };


let count = 0;
setInterval(() => {
  console.log(`Count: ${count++}`);
}, 1000);
