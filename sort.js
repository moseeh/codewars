var min = function(list){
    list.sort((a, b) => a - b)
    return list[0];
}

var max = function(list){
     list.sort((a, b) => b -a )
    return list[0];
}