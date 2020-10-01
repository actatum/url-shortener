import React from "react"
import "./home.styles.scss"
import Shortener from "../components/shortener/shortener.component"

const Home = () => (
    <div className="home">
        <div className="shortener-container">
            <Shortener />
        </div>
    </div>
)

export default Home