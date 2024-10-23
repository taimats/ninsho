import { useAtom } from "jotai"
import { loadingAtom } from "./loadingAtom"

function useJotais() {
    const [loading, setLoading] = useAtom(loadingAtom)
    return {loading, setLoading}
}

export { useJotais }