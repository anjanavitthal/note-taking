# Meeting Notes Example

This is an example application (frontend + backend) for a Markdown Meeting Notes app.


The backend uses an file system to store meeting notes and has three API endpoints: 
* `GET  /note/:id` - Retrieve a note by ID.
* `POST /note` - Create a new note (or update an existing one).
* `GET  /images/:query` - Search for images by using the [Pexels API](https://www.pexels.com/api/).

## Developing locally

## Running

To run the application locally.

```bash
# Run the backend
encore run

# In a different terminal window, run the frontend
cd frontend
npm install
npm run dev
```

### Encore developer dashboard

While `encore run` is running, open [http://localhost:9400/](http://localhost:9400/) to view Encore's local developer dashboard.
Here you can see the request you just made and a view a trace of the response.

## Deployment

### Backend

Deploy your backend to a staging environment in Encore's free development cloud.

```bash
git push encore
```

You can view your backend deploys, metrics and traces at [https://app.encore.dev](https://app.encore.dev).

### Frontend

#### Using GitHub pages

1. Create a repo on GitHub
2. In the `vite.config.js` file, set the `base` property to the name of your repo: 
```ts
base: "/example-meeting-notes/",
```
3. Push your code to GitHub and wait for the GitHub actions workflow to finish.
4. Go to *Settings* → *Pages* for your repo on GitHub and set *Branch* to `gh-pages`.

Your site should now be available at `https://<your-github-username>.github.io/<your-repo-name>/`.

Pushing new code to GitHub will automatically update your site (see the GitHub actions workflow in the `.github` folder).

[Read more about GitHub pages here](https://docs.github.com/en/pages/getting-started-with-github-pages/creating-a-github-pages-site).
