tesseract input.png output -l eng --oem 2 -c tessedit_create_hocr=1 -c tessedit_create_alto=1 -c tessedit_create_txt=1 -c hocr_font_info=0
tesseract --print-parameters

# install tesseract with trained languages from https://github.com/UB-Mannheim/tesseract/wiki
# cli manual https://github.com/tesseract-ocr/tesseract/blob/main/doc/tesseract.1.asc