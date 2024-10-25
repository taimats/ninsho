import axios, { AxiosResponse } from "axios";

async function mockWait(waitTime: number = 1500) {
    return new Promise<void>((resolve) => {
      setTimeout(() => {
        resolve();
      }, waitTime)
    });
  };

async function axiosGet<T>(path: string): Promise<{data: T, err: any}> {
  const r: AxiosResponse = await axios.get(path)
  const resp: { data: T, err: any} = r.data

  await mockWait()
  return resp
}
  
  export { axiosGet, mockWait };