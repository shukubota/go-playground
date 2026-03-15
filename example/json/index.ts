
interface User {
  name: string;
  age: number;
}

const user: User = {
  name: 'Alice',
  age: 30,
}

type BluePrintStatus = 'active' | 'inactive' | 'pending';

interface BluePrintCartAbandonment {
  status: BluePrintStatus;
  detail: {
    delaySeconds: number;
    message: string;
  }
}

interface BluePrintConfirmationOfSafety {
  status: BluePrintStatus;
  detail: {
    safetyLevel: number;
    message: string;
  }
}

const bluePrintCA: BluePrintCartAbandonment = {
  status: 'active',
  detail: {
    delaySeconds: 60,
    message: 'Cart abandonment process is active.',
  },
}

const bluePrintCOS: BluePrintConfirmationOfSafety = {
  status: 'pending',
  detail: {
    safetyLevel: 5,
    message: 'Confirmation of safety is pending.',
  },
}

console.log({ bluePrintCA });
console.log({ bluePrintCOS });
