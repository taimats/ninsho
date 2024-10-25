import { useAtom } from "jotai"
import { testAtom } from "@jotai/testAtom"
import { useJotais } from "@jotai/index"
import { useEffect } from "react"
import { axiosGet } from "@usecase/util"

function Test() {
    const [test] = useAtom(testAtom)
    const { setLoading } = useJotais()

    const load = async () => {
        setLoading(true);
        const data = axiosGet("example")
        console.log("【ログ: data ▶】", data)
        setLoading(false);
    }

    useEffect(() => {
        load();
    }, [])

    return (
        <>
        <p>{test}</p>
        </>
    )
}

export default Test