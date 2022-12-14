export const comment_on_issue = async (id, comment) => {
    await fetch(`https://api.github.com/repos/${GITHUB_USER}/${GITHUB_REPO}/issues/${id}/comments`, {
        body: JSON.stringify({"body": comment}),
        headers: {
            "Accept": "application/vnd.github+json",
            "Authorization": `Bearer ${GITHUB_TOKEN}`,
            "Content-Type": "application/x-www-form-urlencoded",
            "X-Github-Api-Version": "2022-11-28"
        },
        method: "POST"
    })
}

export const close_issue = async (id) => {
    await fetch(`https://api.github.com/repos/${GITHUB_USER}/${GITHUB_REPO}/issues/${id}`, {
        body: JSON.stringify({"state":"closed"}),
        headers: {
            "Accept": "application/vnd.github+json",
            "Authorization": `Bearer ${GITHUB_TOKEN}`,
            "Content-Type": "application/x-www-form-urlencoded",
            "X-Github-Api-Version": "2022-11-28"
        },
        method: "PATCH"
    })
}
