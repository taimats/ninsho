import { useAtom } from "jotai"
import { testAtom } from "@jotai/testAtom"
import { useJotais } from "@jotai/index"
import { useEffect } from "react"
import { mockWait } from "@usecase/util"

function Test() {
    const [test] = useAtom(testAtom)
    const { setLoading } = useJotais()

    const load = async () => {
        setLoading(true);
        await mockWait();
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