import { useJotais } from "@jotai/index"
import { css, cva } from "@styled/css"
import { ReactElement } from "react"

type LoadingProps = { children: ReactElement }

function Loading({ children }: LoadingProps) {
    const { loading } = useJotais()
    return (
        <>
            <div className={css({position: "relative"})}>
                {loading ? 
                <div className={loadingBack()}>
                    <div className={loadingCircle()}></div>
                </div> : null}
                {children}
            </div>
        </>
    )
}

const loadingBack = cva({
    base: {
        backgroundColor: "rgba(128, 128, 128, 0.4)", // rgbaで透明度を指定
        width: "100%",
        height: "100vh",
        position: "absolute",
        zIndex: 100,
    }
})
 
const loadingCircle = cva({
    base: {
      width: "100px",
      height: "100px",
      borderRadius: "50%",
      position: "absolute",
      top: "50%",
      left: "50%",
      border: "20px solid red",
      borderTop: "20px solid transparent",
      animation: "spin 1s linear infinite",
    }
})

export { Loading }