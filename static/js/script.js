let burgerButton = document.querySelector('.burger_button');
let burgerMenu = document.querySelector('.burger_menu');

let activeBurgerMenu = false;

function setActiveBurgerMenu(active) {
  activeBurgerMenu = active;
  if (activeBurgerMenu) {
    burgerButton.classList.add('active');
    burgerMenu.classList.add('active');
  } else {
    burgerButton.classList.remove('active');
    burgerMenu.classList.remove('active');
  }
}
