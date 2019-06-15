using System;

namespace strims.Data 
{
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