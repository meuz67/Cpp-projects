import Hero from './assets/hero.png'
function Card(){
    return (
      <div className="card">
          <img alt="some image" src={Hero} className={"card-image"}></img>
          <h2 className={"card-title"}>Borys Filatov</h2>
          <p className={"card-text"}>Be be be be be</p>
      </div>
    );
}
export default Card