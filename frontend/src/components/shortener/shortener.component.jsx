import React from "react"
import "./shortener.styles.scss"
import axios from "axios"

class Shortener extends React.Component {
    state = {
        url: ''
    }

    handleChange = (e) => {
        const { value, name } = e.target

        this.setState({ [name]: value }, () => console.log(this.state))
    }

    handleSubmit = async (e) => {
        e.preventDefault()
        const { url } = this.state

        try {
            const data = {
                url: url
            }
            const resp = await axios.post("http://localhost:8080/url", data)
            console.log(process.env.REACT_APP_API_URL)
            console.log(resp)
        } catch (err) {
            console.log(err)
        }
    }

    render() {
        return (
            <div className="shortener">
                <h1 className="title">URL Shortener</h1>
                <form onSubmit={this.handleSubmit}>
                    <div className="group">
                        <label className="input-label" htmlFor="url">URL to shorten
                    <input className="input" type="text" name="url" onChange={this.handleChange} />
                        </label>
                        <button className="submit-button">Shorten Me</button>
                    </div>
                </form>
            </div>
        )
    }
}


export default Shortener