let burgerButton = document.querySelector('.burger_button');
let burgerMenu = document.querySelector('.burger_menu');
let body = document.querySelector('body');

let activeBurgerMenu = false;

function scrollToElement(elementId, offset) {
  setActiveBurgerMenu(false);

  var element = document.getElementById(elementId);
  window.scrollTo({
    top: element.offsetTop + offset,
    behavior: 'smooth',
  });
}

function setActiveBurgerMenu(active) {
  if (activeBurgerMenu == active) {
    return;
  }
  activeBurgerMenu = active;
  if (activeBurgerMenu) {
    burgerButton.classList.add('active');
    burgerMenu.classList.add('active');
    body.classList.add('noscroll');
  } else {
    burgerButton.classList.remove('active');
    burgerMenu.classList.remove('active');
    body.classList.remove('noscroll');
  }
}
