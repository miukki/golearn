var myMap = new Map();
myMap.set("0", "foo");
myMap.set(1, "bar");
myMap.set({}, "baz");

for (const [key, value] of myMap.entries()) {
  console.log(key, value);
}


const mapOne = new Map([['a', 1], ['b', 2], ['c' , 3]]);
for (const [key, value] of mapOne.entries()) {
  console.log(key, value);
}

myMap.forEach((value, key) => {
    console.log(value, key)
})