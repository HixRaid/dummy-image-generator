let burgerButton = document.querySelector('.burger_button');
let burgerMenu = document.querySelector('.burger_menu');
let body = document.querySelector('body');

let activeBurgerMenu = false;

let html = document.querySelector('html');

// Smooth scroll to title by id
function scrollToTitle(titleId) {
  setActiveBurgerMenu(false);

  var element = document.getElementById(titleId);
  window.scrollTo({
    top: element.offsetTop - 75,
    behavior: 'smooth',
  });
}

// Switch burger menu
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

function switchTheme() {
  if (html.classList.contains('dark')) {
    html.classList.remove('dark');
    return;
  }

  html.classList.add('dark');
}
