import pytesseract
from PIL import Image

# Load the images into PIL
image_paths = [
    "photo-B80469AD-4562-42CD-87F5-69307DE80180.jpeg",
    "photo-FD940B68-AAF1-4E18-9CF7-D7FAF05560DC.jpeg",
    "photo-7B4B9BF1-744E-40B4-8F76-103609B3A629.jpeg",
    "photo-392D3813-733D-44E4-8476-5592FDFC65CB.jpeg",
    "photo-A283F51B-2252-4417-AA58-571962EF6C87.jpeg",
    "photo-DE2814DC-3D24-4204-8305-35F104DBB1B7.jpeg",
    "photo-1D901109-2994-47CC-A90A-8E8448CB7908.jpeg",
    "photo-6650A0E1-FB19-4AEA-B847-6084ED8EE8F4.jpeg",
    "photo-DC5C8F52-F8EE-4C2D-9747-F41611ED37C3.jpeg",
    "photo-D3FF003A-51AF-4411-BD2A-B29718741314.jpeg",
    "photo-FE2F0441-4CBA-49C7-B99B-C69E5D8567E0.jpeg"
]

output_file_path = 'ocr_output.txt'  # Replace with your desired output file path


# Using Tesseract to do OCR on the images
for image_path in image_paths:
    text = pytesseract.image_to_string(Image.open(image_path))
    # Write the OCR'd text to a file
    with open(output_file_path, 'a') as file:
        file.write(text)

