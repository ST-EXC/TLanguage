let a = 1;
while(a < 5){
    println("a=",a);
    a = a + 1;
}

let lessThan = fn(a,b){
    return a<b;
}
println("=================");
while(lessThan(a,10)){
    println("a=",a);
    a = a + 1;
}
