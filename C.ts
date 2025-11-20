/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-20
 * @fileoverview If statement problem.
 */

// input
const NAME: string = prompt("What is your name? \n") || ("n/a");

const ageString: string = prompt("How old are you? \n") || ("0");
const ageNumber: number = parseInt(ageString);

const STUDENT: string = prompt("Are you a student? \n") || ("no");

// processing

if (STUDENT.toLowerCase() == "yes") {
  if (ageNumber >= 13 && ageNumber <= 19) {
    console.log(`Student ${NAME} is a teenager. \n`);
  } else if (ageNumber >= 5 && ageNumber <= 12) {
    console.log(`Student ${NAME} is a child. \n`);
  } else {
    console.log(`${NAME} is in a different life stage. \n`);
  }
} else if (STUDENT.toLowerCase() == "no") {
  if (ageNumber >= 20 && ageNumber <= 30) {
    console.log(`${NAME} is a young adult.`);
  } else {
    console.log(`${NAME} is in a different life stage.`);
  }
}

console.log("\nDone.");
