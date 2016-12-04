//var w=window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
var theWidth = 0;
function setWidth(){
	window.alert("Function setWidth has been ran.");
	//var theWidth = 0;
    var w = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;

            if (w > 1024) {
                theWidth = 1024;
            } else {
                theWidth = w;
            }
    //return theWidth;
    var myElement = document.querySelector("#wrapper");
    myElement.style.width = theWidth;
    
}

//window.onload = setWidth();       // When the page first loads
//window.onresize = setWidth();     // When the browser changes size
//document.getElementById("wrapper").style.width = theWidth - 900;

//window.alert("Started bonemaster.js");
function displaytheWidthNow(){
	var w = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
	window.alert("The Width is: " + w);
	window.alert("theWidth is: " + theWidth);
}
