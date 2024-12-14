let fab = fn(x){
    if(x < 2){
        return 1;
    }else{
        return fab(x-2)+fab(x-1);
    }
}
println(fab(10));

let x = 1;
let y = 2;

let add = fn(x,y){
    return x+y;
}

println("add(x,y)=",add(x,y));
println("add(x,add(x,2))=",add(x,add(x,2)));

let add2 = fn(){
    let c = 2;
    return fn(x,y){
        return x+y+c;
    };
}

println("add2()=",add2())
println("add2()(x,y)=",add2()(x,y));

let add3 = fn(add,x,y){
    return add(x,y);
}

println("add3(add,x,y)=",add3(add,x,y));
