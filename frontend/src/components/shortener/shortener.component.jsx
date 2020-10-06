import React from "react"
import "./shortener.styles.scss"
import axios from "axios"
import { Link } from "react-router-dom"

class Shortener extends React.Component {
  state = {
    url: '',
    top5: null,
  }

  async componentDidMount() {
    try {
      const resp = await axios.get(`${process.env.REACT_APP_API_URL}/url`)
      this.setState({ top5: resp.data })
    } catch (err) {
      console.log(err)
    }
  }

  handleSlug = (e, url) => {
    console.log(e.target)
    console.log(url)
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
      await axios.post(`${process.env.REACT_APP_API_URL}/url`, data)
    } catch (err) {
      console.log(err)
    }
  }

  render() {
    const { top5 } = this.state
    let listItems = null
    console.log(top5)
    if (top5 === null) {
      console.log("top 5 is empty")
    } else {
      listItems = top5.map((el) => <Link to={el.slug} > <div className="top5item" key={el.slug}><p>slug: {el.slug} url: {el.url} uses: {el.uses}</p></div></Link >)
    }
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
        <h1>Top 5 Slugs</h1>
        {top5 !== null ? listItems : null}
      </div>
    )
  }
}

export default Shortener