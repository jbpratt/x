using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;
using System.Linq;
using Microsoft.AspNetCore.Blazor;

namespace strims.Data
{
  public class AppState
  {
    public IReadOnlyList<Strim> Strims {get; private set;}
    
    private readonly HttpClient http;
    public AppState(HttpClient httpInstance)
    {
      http = httpInstance;
    }

    public async Task LoadStrims()
    {
      if(Strims == null)
      {
        Strims = await http.GetJsonAsync<Strim[]>("strims.json");
      }
    }

    public Strim LoadStrim(int strimId)
    {
      return Strims.SingleOrDefault(strim => strim.id == strimId);
    }
  }
  public class Strim 
    {
        public bool Live { get; set; }
        public bool Nsfw { get; set; }
        public bool Hidden { get; set; }
        public bool Afk { get; set; }
        public Int16 Rustlers { get; set; }
        public Int16 AfkRustlers { get; set; }
        public String Service { get; set; }
        public String Channel { get; set; }
        public String Title { get; set; }
        public String Thumbnail { get; set; }
        public String Url { get; set; }
        public Int16 Viewers { get; set; }
    }
}
