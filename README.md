# Plug or Plugs, idk

## About Plug/Plugs - Motivation

After seeing the low amount of well-defined online services offering user-to-user help, I looked into creating a solution myself.
One of the most important things was for it to be quick, easy to use and trustworthty for the user-to-user.
Plug/Plugs will try to deliver a system where you access help from anyone, at anytime. This concept is called "Plugs".
Plugs will also have the option to list what they're capable of doing, and the time they're available (in realtime)

On the other hand, are jobs. These are your traditonal jobs which users can post, and anyone can offer to help in exchange for $$$.

Clearly, the "Plugs" will the most interesting concecpt of the application, and how to perfect it's use case. It'll be an exciting project which I hope
to keep open-source. (plz don't steal my idea. Contribute instead)

## Technology and project details

This project is a monorepo, and split up in three different apps/packages.

- `frontend` - The React application for Plugs/Plugs.
- `api` - The backend built in Go.
- `packages` - Other reusable packages used internally within the React projects. Such as configs, ui libs and more.
- `native` - (CTA) The react-native (expo) applicaiton for Android and iOS.

### Frontend / App

I decided to for a more traditional architecture for this projects due to it's potential scale. I considered a fullstack framework like Next.js or Remix.run, but I think
this would be a mess to maintain in the long-term. There are trade-off's with the current architecture, one being SEO, which I actually rather concered about :).
Other than what's already mentioned, I believe this is a more sustainable approach for the frontend.

The choice for using Expo, instead of just React-Native, is the ease of use, and abstractions I don't want to spend time on handling myself. Building the app Expo is also awesome.

### Backend

Nothing much here. I chose Go because I wanted a fast and robsut backend, something Node does not offer :eyes:

Seriously, I like Go, as well as it would be nice to learn more. Would probably have gone for Node if I wanted to it quick.

## Funding

I need investors!
