/*
 *	JSLists v0.3.7
 *	© 2016 George Duff
 * 	Release date: 01/06/2016
 *	The MIT License (MIT)
 *	Copyright (c) 2016 George Duff
 *	Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *	The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */
(function(){
	'use strict';
	function define_JSLists(){
		//********************
		//Private variables
		//********************
		var JSLists = {};

		//********************
		//Private methods
		//********************
		var JSLists_Error = function(error, alertType){
			console.log(error);
		}
		var getUl = function(){
			return document.getElementsByTagName("UL");
		};

		var getOl = function(){
			return document.getElementsByTagName("OL");
		};

		var getAllLists = function(){
			var olLists = Array.prototype.slice.call(document.getElementsByTagName("UL")),
				ulLists = Array.prototype.slice.call(document.getElementsByTagName("OL"))
			var gLists = olLists.concat(ulLists);
			return gLists;
		}
		
		var searchList = function(listId, searchTerm){
			var olLists = Array.prototype.slice.call(document.getElementsByTagName("UL")),
				ulLists = Array.prototype.slice.call(document.getElementsByTagName("OL")),
				liItems = Array.prototype.slice.call(document.getElementsByTagName("LI"))
			var gLists = olLists.concat(ulLists);
			return gLists;
		}
		//********************
		//Public variables
		//********************

		//********************
		//Public methods
		//********************
		JSLists.greet = function(){
			console.log("** Welcome to JSLists **");
		};

		JSLists.checkNodes = function(){
		};

		JSLists.collapseAll = function(listId){
			var i, olLists = Array.prototype.slice.call(document.getElementsByTagName("UL")),
				ulLists = Array.prototype.slice.call(document.getElementsByTagName("OL"))
			var gLists = olLists.concat(ulLists);

			for(i=1; i<gLists.length; i++){
				gLists[i].setAttribute('class', 'jsl-collapsed');
			};
		};

		JSLists.openAll = function(listId){
			var i, olLists = Array.prototype.slice.call(document.getElementsByTagName("UL")),
				ulLists = Array.prototype.slice.call(document.getElementsByTagName("OL"))
			var gLists = olLists.concat(ulLists);

			for(i=1; i<gLists.length; i++){
				gLists[i].setAttribute('class', 'jsl-open');
			};
		};

		JSLists.generateCss = function(){
			var css = document.createElement('style'); //Should all this be a seperate CSS file?
			var styles = ".jslist-li{margin-left: 22px; height: 24px;}";	//This is for the LI with a list below
			styles += ".jslist-ul{margin-left: 22px;}";	//This is for all UL's
			styles += ".jslist-ol{margin-left: 22px;}";	//This is for all OL's
			styles += ".jsl-collapsed{display: none;}";
			styles += ".jsl-collapsed-arrow{float: left; clear: both; margin-right: 11px; width: 11px; height: 11px; cursor: pointer; background: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAsAAAALCAYAAACprHcmAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAKJJREFUeNqEkc0JxCAQhUfRg7EJweN2IikoJWxBkm4EexCvur4gi4kuO/Dw532O48i89y8ieje5Jk5zlKaz6RAArbW7MYYYYxNZa+Uxxj2EcGVyALsxCdF9B5gj4wgopW7rfiNf1fgzxFAbaa2/xrZt15hznuHRwKERWmZ+dGEJF7RnNFNKN7jPCx54tj6SEIKklJOwDx8cMh9oeNPfH/wIMABbu2PPHYUsbQAAAABJRU5ErkJggg==') no-repeat;}";
			styles += ".jsl-open{display: block;}";
			styles += ".jsl-open-arrow{background: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAsAAAALBAMAAABbgmoVAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAACRQTFRFtLS0tLS0tLS0AAAAtLS019fX8/Pz9PT09fX19vb29/f3+Pj4NWr6kwAAAAN0Uk5TK7P9wooeBQAAAD9JREFUCNdjYFRxcVFkYBBdvXv3UgEGrVVAoMDgPXPmzNkOcKrZ2LgbSHV0dICo8vLyagcGrbS0tCwFmD6IKQDtKxlF/vrVqgAAAABJRU5ErkJggg==') no-repeat;}";
			if (css.styleSheet) css.styleSheet.cssText = styles;
			else css.appendChild(document.createTextNode(styles));
			document.getElementsByTagName("head")[0].appendChild(css);
		};

		JSLists.paddULLists = function(listId){
			var i, listItems = document.getElementById(listId).getElementsByTagName("UL");
			for(i=0; i<listItems.length; i++){
				listItems[i].classList.add('jslist-ul');
			}
		};
		JSLists.paddOLLists = function(listId){
			var i, listItems = document.getElementById(listId).getElementsByTagName("UL");
			for(i=0; i<listItems.length; i++){
				listItems[i].classList.add('jslist-ol');
			}
		};

		JSLists.padLists = function(listId){
			var i, listItems = document.getElementById(listId).getElementsByTagName("LI");
			for(i=0; i<listItems.length; i++){
				if(listItems[i].childNodes[0].className != "jsl-collapsed-arrow"){
					listItems[i].classList.add('jslist-li');
				}
			}
			this.paddULLists(listId);
			this.paddOLLists(listId);
		};

		JSLists.generateList = function(listId){
			this.generateCss();
			document.getElementById(listId).style.display = "none;"
			var i, j, curElem, ulCount, listItems = document.getElementById(listId).getElementsByTagName('LI'); //this should be the main parent
			for(i=0; i<listItems.length; i++){
				if(listItems[i].id.length > 0){ //if node already has an id
					curElem = document.getElementById(listItems[i].id);
					ulCount = document.getElementById(listItems[i].id).getElementsByTagName("UL");
					if(ulCount.length > 0){ //There is a nested UL in this LI element, now find the position of the UL
						for(j=0; j<ulCount.length; j++){
							if(ulCount[j].nodeName == "UL" || ulCount[j].nodeName == "OL"){ //** Or add in OL
								break; //Multiple UL's? //Set class collapseAll here
							}
						}
						ulCount[j].setAttribute('class', 'jsl-collapsed');
						//Now make the div and insert as first node
						var tglDiv = document.createElement("div");
						tglDiv.setAttribute('class', 'jsl-collapsed-arrow');
						tglDiv.setAttribute("id", listItems[i].id + i +'_tgl');
						curElem.insertBefore(tglDiv, curElem.childNodes[0]);
						document.getElementById(listItems[i].id + i +'_tgl').addEventListener('click', function(e){
							document.getElementById(e.target.id).classList.toggle('jsl-open-arrow');
							document.getElementById(e.target.id).parentElement.lastElementChild.classList.toggle('jsl-open');
							e.stopPropagation();
						},true);
					}
				}else{
					//Add id to the node
					listItems[i].setAttribute("id", listId+"tmp"+i);
					curElem = document.getElementById(listId+"tmp"+i);
					ulCount = document.getElementById(listItems[i].id).getElementsByTagName("UL");

					if(ulCount.length > 0){ //There is a nested UL in this LI element, now find the position of the UL
						for(j=0; j<ulCount.length; j++){
							if(ulCount[j].nodeName == "UL" || ulCount[j].nodeName == "OL"){
								break; //Multiple UL's? //Set class collapseAll here
							}
						}
						ulCount[j].setAttribute('class', 'jsl-collapsed');
						var tglDiv = document.createElement("div");
						tglDiv.setAttribute('class', 'jsl-collapsed-arrow');
						tglDiv.setAttribute("id", listItems[i].id + i +'_tgl');
						curElem.insertBefore(tglDiv, curElem.childNodes[0]);
						document.getElementById(listItems[i].id + i +'_tgl').addEventListener('click', function(e){
							document.getElementById(e.target.id).classList.toggle('jsl-open-arrow');
							document.getElementById(e.target.id).parentElement.lastElementChild.classList.toggle('jsl-open');
							e.stopPropagation();
						},true);
					}
					listItems[i].removeAttribute("id");
				}
			}
			setTimeout(function(){
				document.getElementById(listId).style.display = "block;"
			},99);

			this.padLists(listId);
		};

		JSLists.applyToList = function(listId, listType){
			// if(typeof listType === undefined || listType.toUpperCase() != "UL" || listType.toUpperCase() != "OL"){
				// console.error('ERROR:\nInvalid list type'); //Still to do back out if no valid type
			// }
			switch(listType.toUpperCase()){
				case "UL":
					this.generateList(listId, "UL");
					break;
				case "OL":
					this.generateList(listId, "OL");
					break;
				case "ALL":
					this.generateList(listId);
					break;
				default:
					break;
			}
		};
	return JSLists;
}
	//define the JSLists library globally if it doesn't already exist
	if(typeof(JSLists) === 'undefined'){
		window.JSLists = define_JSLists();
	}else{
		console.log("JSLists already defined.");
	}
})();