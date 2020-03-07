## Yelp Help v2

### What is this?
In the winter of 2018 I wrote a small web app with Express and React for my portfolio. The idea was simple, it would return a random restaurant close to the user to solve the problem of not being able to pick a place to eat. After a year of working at Beam Dental as a full stack software engineer, I decided to look back at the code I had written for this small app. I was actually pretty happy with what I had, but there were a few things that jumped out right away that I totally missed the mark on.

### Tech Stack
- Golang
- React

### Running the Project
```bash
docker build -t yelp-help-v2-server .
docker run -d -p 8000:8000 yelp-help-v2-server
```

### Testing the Project
TBD