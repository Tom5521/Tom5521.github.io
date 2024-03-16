const imgs = [
  "./images/zazu/zazu.jpg",
  "./images/zazu/born-to-dilly-dally.jpg",
  "./images/zazu/theres-motion.jpg",
  "./images/zazu/zazu-in-the-war.jpg",
  "./images/zazu/zazu-in-watermelon.jpg",
];
const number = Math.floor(Math.random() * imgs.length);
document.getElementById("zazu-img").src = imgs[number];
