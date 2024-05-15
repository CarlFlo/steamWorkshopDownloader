document.addEventListener("DOMContentLoaded", function() {
    const inputField = document.getElementById("inputText");
    const submitButton = document.querySelector("button");
    const responseDiv = document.getElementById("response");

    const urlRegex = /^https:\/\/steamcommunity\.com\/sharedfiles\/filedetails\/\?id=\d+$/;
    let debounceTimeout;

    inputField.addEventListener("input", handleInputChange);
    document.getElementById("myForm").addEventListener("submit", handleSubmit);

    function handleInputChange() {
        clearTimeout(debounceTimeout);
        debounceTimeout = setTimeout(() => {
            const inputValue = inputField.value;
            if (urlRegex.test(inputValue)) {
                verifyInput(inputValue);
            } else {
                submitButton.disabled = true;
                displayWorkstopItemInformation()
            }
        }, 400); // 400ms debounce time
    }

    function verifyInput(inputValue) {
        const localStorageKey = inputValue.split("=")[1];
        const cachedResult = localStorage.getItem(localStorageKey);
        
        if (cachedResult) {
            try {
                const data = JSON.parse(cachedResult);
                
                displayWorkstopItemInformation(data);
                updateSubmitButton(true);

            } catch (error) {
                console.error("Parsing failed:", error);
                updateSubmitButton(false);
            }
        } else {
            fetchInputData(inputValue, localStorageKey);
        }
    }

    function displayWorkstopItemInformation(data) {
        
        const objectDisplay = document.getElementById('objectDisplay');

        if (!data) {
            // Hide the object display section if there is no data
            objectDisplay.style.display = 'none';
        }

        console.log(data)

        const itemData = {
            name: data.item_name,
            lastUpdated: data.last_updated,
            creator: "temp",
            creatorLink: data.created_by,
            fileSize: data.file_size,
            previewImage: data.preview_image
        };
    
        // Populate the object display section
        document.getElementById('itemName').textContent = itemData.name;
        document.getElementById('lastUpdated').textContent = `Last Updated: ${itemData.lastUpdated}`;
        document.getElementById('creator').innerHTML = `Creator: <a id="creatorLink" href="${itemData.creatorLink}" target="_blank">${itemData.creator}</a>`;
        document.getElementById('fileSize').textContent = `File Size: ${itemData.fileSize}`;
        document.getElementById('previewImage').src = itemData.previewImage;

        // Show the object display section
        objectDisplay.style.display = 'block';
       
    }

    function fetchInputData(inputValue, localStorageKey) {
        fetch("/getItemInfo", {
            method: "POST",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            body: "inputText=" + encodeURIComponent(inputValue),
        })
        .then(response => {
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            return response.json(); // assuming the response is in JSON format
        })
        .then(data => {
            displayWorkstopItemInformation(data);
            localStorage.setItem(localStorageKey, JSON.stringify(data));
            updateSubmitButton(true);
        })
        .catch(error => {
            console.error("Error:", error);
            updateSubmitButton(false);
        });
    }

    function updateSubmitButton(isValid) {
        submitButton.disabled = !isValid;
    }

    function handleSubmit(event) {
        event.preventDefault();

        const inputText = inputField.value;
        disableSubmitButton("Loading...");

        fetch("/submit", {
            method: "POST",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            body: "inputText=" + encodeURIComponent(inputText),
        })
        .then(handleFetchResponse)
        .then(handleFileDownload)
        .catch(handleError)
        .finally(resetSubmitButton);
    }

    function disableSubmitButton(text) {
        submitButton.disabled = true;
        submitButton.innerText = text;
    }

    function resetSubmitButton() {
        submitButton.disabled = false;
        submitButton.innerText = "Download";
    }

    function handleFetchResponse(response) {
        if (!response.ok) {
            throw new Error("Network response was not ok");
        }
        return response.blob().then(blob => {
            const disposition = response.headers.get('Content-Disposition');
            let filename = "download.zip";
            if (disposition && disposition.indexOf('attachment') !== -1) {
                const match = disposition.match(/filename="(.+)"/);
                if (match && match[1]) {
                    filename = match[1];
                }
            }
            return { blob, filename };
        });
    }

    function handleFileDownload({ blob, filename }) {
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.style.display = 'none';
        a.href = url;
        a.download = filename;
        document.body.appendChild(a);
        a.click();
        window.URL.revokeObjectURL(url);
    }

    function handleError(error) {
        console.error("Error:", error);
        responseDiv.innerText = "Error: " + error.message;
    }
});
