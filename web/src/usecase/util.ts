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

async function axiosDelete<T>(path: string): Promise<{data: T, err: any}> {
  const r: AxiosResponse = await axios.delete(path)
  const resp: { data: T, err: any} = r.data

  await mockWait()
  return resp
}

async function axiosPost<T>(path:string, payload: any): Promise<{data: T, err: any}> {
  const params = JSON.stringify(payload)
  const r = await axios.post(path, params)
  const resp: {data: T, err: any} = r.data

  await mockWait()
  return resp
}
  
export { axiosGet, axiosDelete, axiosPost, mockWait };