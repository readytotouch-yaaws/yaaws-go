// Code generated by qtc from "online.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package v1

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamOnline(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Yet another anonymous work search</title>
    <meta name="description" content="Yet another anonymous work search">
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	<link rel="author" type="text/plain" href="https://readytotouch.com/humans.txt"/>

	`)
	streamfavicon(qw422016)
	qw422016.N().S(`

	<link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500&display=swap" rel="stylesheet">
    <script src="https://cdn.jsdelivr.net/npm/apexcharts"></script>

    `)
	streamonlineStyles(qw422016)
	qw422016.N().S(`
</head>
<body>
`)
	streamheader(qw422016)
	qw422016.N().S(`
<main class="main-wrapper">
<section class="hero">
    <div class="hero__wrapper">
        <div class="hero__header">
            <div class="hero__status">
                Service status: <span class="hero__progress">Work in progress</span>
            </div>
            <a href="https://readytotouch.com/" target="_blank" class="hero__link">
                <span class="hero__link-title">readytotouch.com</span>
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M7 17L17 7M7 7h10v10" stroke="#C6CAD2" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
            </a>
        </div>
        <div class="hero__main">
            <h1 class="hero__title">Anonymous job search</h1>
            <p class="hero__subtitle">Tracking registrations and online users of the platform. Sign in with GitHub or GitLab to support the project</p>
        </div>
        <div class="hero__buttons-group">
            <a href="#" class="hero__button">
                <img src="/assets/images/pages/online/github-white.svg" alt="GitHub" class="button-icon">Log in with GitHub
            </a>
            <a href="#" class="hero__button">
                <img src="/assets/images/pages/online/gitlab.png" alt="GitLab" class="hero__button-icon">Log in with GitLab
            </a>
        </div>
    </div>
</section>
<div class="container">
    `)
	streamstats(qw422016)
	qw422016.N().S(`
    `)
	streamregistrationHistory(qw422016)
	qw422016.N().S(`
    `)
	streamprojects(qw422016)
	qw422016.N().S(`
    `)
	streamarticles(qw422016)
	qw422016.N().S(`
</div>
</main>
`)
	streamfooter(qw422016)
	qw422016.N().S(`
</body>
</html>
`)
}

func WriteOnline(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamOnline(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func Online() string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteOnline(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
