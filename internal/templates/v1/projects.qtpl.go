// Code generated by qtc from "projects.qtpl". DO NOT EDIT.
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

func streamprojects(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
<section class="projects">
    <h2 class="projects__title">Other projects</h2>
    <div class="projects__wrapper">
        <div class="projects__item project big-card">
            <div class="project__info-card">
                <div class="project__stars">
                    <iframe src="https://ghbtns.com/github-btn.html?user=u8views&repo=go-u8views&type=star&count=true&size=large"
                            frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
                </div>
                <div class="project__info-group">
                    <h3 class="project__title">u8views — views counter</h3>
                    <p class="project__subtitle">GitHub profile views statistics</p>
                </div>
                <a href="https://u8views.com/" target="_blank" class="project__link">u8views.com</a>
            </div>
            <div class="project__view">
                <img src="/assets/images/pages/online/u8views.png" class="project__view-img" alt="project view">
            </div>
        </div>
        <div class="projects__item project big-card">
            <div class="project__info-card">
                <div class="project__stars">
                    <iframe src="https://ghbtns.com/github-btn.html?user=doutivity&repo=doutivity.github.io&type=star&count=true&size=large"
                            frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
                </div>
                <div class="project__info-group">
                    <h3 class="project__title">doutivity</h3>
                    <p class="project__subtitle">Search for a comment among comments published on DOU</p>
                </div>
                <a href="https://doutivity.github.io/" target="_blank"
                   class="project__link">doutivity.github.io</a>
            </div>
            <div class="project__view">
                <img src="/assets/images/pages/online/doutivity.png" class="project__view-img" alt="project view">
            </div>
        </div>
        <div class="projects__group">
            <div class="projects__item project small-card">
                <div class="project__info-card">
                    <div class="project__stars">
                        <iframe src="https://ghbtns.com/github-btn.html?user=json-to-proto&repo=json-to-proto.github.io&type=star&count=true&size=large"
                                frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
                    </div>
                    <div class="project__info-group">
                        <h3 class="project__title">JSON-to-Proto</h3>
                        <p class="project__subtitle">Convert JSON to Protobuf</p>
                    </div>
                    <a href="https://json-to-proto.github.io/" target="_blank"
                       class="project__link">json-to-proto.github.io</a>
                </div>
                <div class="project__view">
                    <img src="/assets/images/pages/online/open-project.png" class="project__view-img"
                         alt="project view">
                </div>
            </div>
            <div class="projects__item project small-card">
                <div class="project__info-card">
                    <div class="project__stars">
                        <iframe src="https://ghbtns.com/github-btn.html?user=xml-to-go&repo=xml-to-go.github.io&type=star&count=true&size=large"
                                frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
                    </div>
                    <div class="project__info-group">
                        <h3 class="project__title">XML-to-Go</h3>
                        <p class="project__subtitle">Convert XML to Go</p>
                    </div>
                    <a href="https://xml-to-go.github.io/" target="_blank"
                       class="project__link">xml-to-go.github.io</a>
                </div>
                <div class="project__view">
                    <img src="/assets/images/pages/online/open-project.png" class="project__view-img"
                         alt="project view">
                </div>
            </div>
        </div>
    </div>
</section>
`)
}

func writeprojects(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamprojects(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func projects() string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeprojects(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
