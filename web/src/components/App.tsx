import { Link, Routes, Route } from 'react-router-dom'
import Test from '@components/test/Test'
import { Loading } from '@components/UI/Loading'
import { css } from '../styled-system/css'

function App() {
  return (
    <>
    <Loading>
      <div>
        <div className={css({ fontSize: "2xl", fontWeight: "bold", color: "red" })}>Hello Red!</div>
        <Link to="/test"><h1>Test</h1></Link>

        <Routes>
          <Route path="/test" element={<Test />} />
        </Routes>
      </div>
    </Loading>
    </>
  )
}

export default App
