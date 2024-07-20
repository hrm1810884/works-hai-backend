import sys
import os
from PIL import Image

def parse_args(args):
    image_dict = {}
    for arg in args:
        key, value = arg.split("=")
        image_dict[key] = value
    return image_dict

def process_image(image_paths, out_path):
    print("python called")
    for image_path in image_paths:
        try:
            with Image.open(image_path) as img:
                # Process the image and save it
                img.save(out_path)
                print(f"Processed image saved as {out_path}")
        except Exception as e:
            print(f"Failed to process {image_path}: {e}")

if __name__ == "__main__":
    if len(sys.argv) < 3:
        print("Usage: python3 process_image.py <key=value> <key=value> ... <output_file_path>")
        sys.exit(1)

    image_dict = parse_args(sys.argv[1:len(sys.argv)-1])
    out_path = sys.argv[len(sys.argv)-1]
    image_paths = list(image_dict.values())
    process_image(image_paths, out_path)
