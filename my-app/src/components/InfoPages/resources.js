
export const Resources = ({ resource }) => {
    return (
      <div>
        <center><h1>Resource List</h1></center>
        {Resources.map((resource) => (
          <div class="card">
            <div class="card-body">
              <h5 class="title">{resource.title}</h5>
              <h6 class="date">{resource.date}</h6>
              <h6 class="description">{resource.description}</h6>
              <h6 class="author">{resource.author}</h6>
              <h6 class="type">{resource.type}</h6>
            </div>
          </div>
        ))}
      </div>
    )
  };
  
  export default Resources;