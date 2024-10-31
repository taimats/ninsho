import { useJotais } from "@jotai/index"
import { SyntheticEvent, useEffect, useState } from "react"
import { axiosGet, axiosDelete, axiosPost } from "@usecase/util"
import { css } from "@styled/css"

type Example = {id: number, name: string}
type Form = {name: string}
const initForm: Form = {name: ""}

type FormProps = {
    complete: () => Promise<void>
}

const Form = (p: FormProps) => {
    const {setLoading} = useJotais()
    const [form, setForm] = useState<Form>(initForm)

    const create = async (e: SyntheticEvent) => {
        e.preventDefault()
        setLoading(true)

        try {
            await axiosPost("example/insert", form)
            await p.complete()
        } catch (error: any) {
            alert(error.message)
        } finally {
            setForm(initForm)
            setLoading(false)
        }
    }

    const onChange = (e: SyntheticEvent) => {
        const target = e.target as HTMLInputElement
        setForm({name: target.value})
    }

    return (
        <form method="POST" onSubmit={create}>
            <input className={inputForm} type="text" value={form.name} onChange={onChange} />
            <button type="submit">登録</button>
        </form>
    )
}

function Test() {
    const [examples, setExamples] = useState<Example[]>([])
    const { setLoading } = useJotais()

    const load = async () => {
        setLoading(true);
        try {
            const resp = await axiosGet<Example[]>("example/all")
            setExamples(resp.data) 
        } catch (error: any) {
            alert(error.response.data.err)
        } finally {
            setLoading(false);
        }
    }

    const del = async (id:number) => {
        setLoading(true)
        const path = `example/delete/${id}`
        try {
            await axiosDelete(path)
            await load()
        } catch(error: any) {
            alert(error.response.data.err)
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        load();
    }, [])

    return (
        <>
        <ul>
            {examples.map((e: Example) => {
                return (
                    <li key={e.id} className={card} onClick={() => { del(e.id) }}>
                        {e.name}
                    </li>
                )
            })}
        </ul>
        <Form complete={load}/>
        </>
    )
}

const card = css({
    backgroundColor: "red",
    color: "white",
    margin: "10px",
    padding: "10px",
    width: "200px",
})

const inputForm = css({
    border: "1px solid #000",
    margin: "10px",
});

export default Test