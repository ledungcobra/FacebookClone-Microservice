export const extractErrorFromResponse = (response) => {
    console.log(response)
    if (!response) {
        return ""
    }
    const errors = response.errors;
    if (!errors) {
        return ""
    }
    return response.message + "\n" +
        errors.map(e => "Field " + e.field + ': ' + e.errors.join(' ')).join('\n')
}