import React from "react"
import "./shortener.styles.scss"
import axios from "axios"

class Shortener extends React.Component {
    state = {
        url: '',
        top5: [],
    }

    async componentDidMount() {
        try {
            const resp = await axios.get("http://localhost:8080/url")
            this.setState({ top5: resp.data })
        } catch (err) {
            console.log(err)
        }
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
        } catch (err) {
            console.log(err)
        }
    }

    render() {
        const { top5 } = this.state
        console.log(top5)
        // const listItems = top5.map((el) => <li key={el.slug}>{el}</li>)
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

                <div className="top5">
                    <ul>

                    </ul>
                </div>
            </div>
        )
    }
}


export default Shortener