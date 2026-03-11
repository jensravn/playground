import argparse
import json
import os
import random

from PIL import Image, ImageDraw


def get_color(label):
    """Generate a random, but consistent color based on the label name."""
    random.seed(label)
    return (random.randint(50, 255), random.randint(50, 255), random.randint(50, 255))


def main():
    # Set up command line arguments
    parser = argparse.ArgumentParser(
        description="Draw bounding boxes from a JSON file onto an image."
    )
    parser.add_argument(
        "image", nargs="?", default="./image.jpg", help="Path to the input image file"
    )
    parser.add_argument(
        "json", nargs="?", default="./data.json", help="Path to the JSON file containing the data"
    )
    args = parser.parse_args()

    # 1. Load the JSON data
    try:
        with open(args.json, "r") as f:
            data = json.load(f)
    except Exception as e:
        print(f"Error reading JSON file: {e}")
        return

    # 2. Load the image
    try:
        img = Image.open(args.image).convert("RGB")
    except Exception as e:
        print(f"Error reading Image file: {e}")
        return

    draw = ImageDraw.Draw(img)
    width, height = img.size

    # 3. Draw the boxes
    for item in data:
        box = item.get("box_2d")
        label = item.get("label", "unknown")

        if not box or len(box) != 4:
            continue

        # Convert AI [ymin, xmin, ymax, xmax] coordinates (scaled 0-1000) to actual image pixels
        ymin = (box[0] / 1000.0) * height
        xmin = (box[1] / 1000.0) * width
        ymax = (box[2] / 1000.0) * height
        xmax = (box[3] / 1000.0) * width

        # Get a unique color for this specific label
        color = get_color(label)

        # Draw the bounding box
        draw.rectangle([xmin, ymin, xmax, ymax], outline=color, width=3)

        # Draw the text label just above the box
        text_y = ymin - 15 if ymin > 15 else ymin + 5
        draw.text((xmin, text_y), label, fill=color)

    # 4. Save the resulting image dynamically
    base_name, ext = os.path.splitext(args.image)
    output_filename = f"{base_name}_annotated{ext}"

    img.save(output_filename)
    print(f"Success! Annotated image saved as: {output_filename}")
    img.show()


if __name__ == "__main__":
    main()
