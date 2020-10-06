import React from "react"
import axios from "axios"
import { Link } from "react-router-dom"

class Slug extends React.Component {

  state = {
    url: '',
    errorMessage: ''
  }

  async componentDidMount() {
    try {
      const slug = this.props.location.pathname.slice(1)
      const resp = await axios.get(`${process.env.REACT_APP_API_URL}/url/${slug}`)
      this.setState({ url: resp.data.url })
    } catch (err) {
      this.setState({ errorMessage: err.response.data.message })
    }
  }

  render() {
    const { url, errorMessage } = this.state
    return (
      <div>
        {url !== '' ? (window.location = url) :
          <div>
            <h3>{errorMessage}</h3>
            <Link to="/"><button>Go back</button></Link>
          </div>
        }
      </div>
    )
  }
}

export default Slug