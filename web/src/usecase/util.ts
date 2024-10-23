async function mockWait(waitTime: number = 1500) {
    return new Promise<void>((resolve) => {
      setTimeout(() => {
        resolve();
      }, waitTime)
    });
  };
  
  export { mockWait };