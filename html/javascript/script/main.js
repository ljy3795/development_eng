const myHeading = document.querySelector('h1');
// myHeading.textContent = 'Hello world!';
myHeading.textContent = multiply(3,4);

let iceCream = 'chocolate';
if(iceCream === 'chocolate') {
  alert('Yay, I love chocolate ice cream!');    
} else {
  alert('Awwww, but chocolate is my favorite...');    
}

function multiply(num1,num2) {
    let result = num1 * num2;
    return result;
}

// document.querySelector('html').onclick = function() {
//     alert('Ouch! Stop poking me!');
// }

let myHTML = document.querySelector('html');
myHTML.onclick = function() {
    alert('Ouch! Stop poking me!');
};


let myImage = document.querySelector('img');

myImage.onclick = function() {
    let mySrc = myImage.getAttribute('src');
    if(mySrc === 'images/firefox-icon.png') {
      myImage.setAttribute ('src','images/firefox2.png');
    } else {
      myImage.setAttribute ('src','images/firefox-icon.png');
    }
}