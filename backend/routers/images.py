from typing import Annotated

from fastapi import APIRouter, HTTPException, UploadFile, File

router = APIRouter(
    prefix="/images",
    tags=["images"],
    responses={404: {"description": "Not found"}},
)

fake_images = [{"image_id": "foo123", "filename": "dog1.jpg"}, {"image_id": "baz456", "filename": "deebo.png"}]

def find(lst: list, key: str, value: str) -> int | None:
    for i, d in enumerate(lst):
        if d[key] == value:
            return i
    return None

@router.get("/")
async def get_images():
    return fake_images

@router.post("/")
async def upload_image(file: Annotated[UploadFile, File()]):
    return {"filename": file.filename, "size": file.size, "type": file.content_type}

@router.get("/{image_id}")
async def get_image(image_id: str):
    if not find(fake_images, "image_id", image_id):
        raise HTTPException(status_code=404, detail="Image not found")
    return dict(image for image in fake_images if image["image_id"] == image_id)

@router.delete("/{image_id}", status_code=204)
async def delete_image(image_id: str):
    idx = find(fake_images, "image_id", image_id)
    if not idx:
        raise HTTPException(status_code=404, detail="Image not found")
    del fake_images[idx]
    return "deleted"
