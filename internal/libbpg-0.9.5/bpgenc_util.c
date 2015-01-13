
#include "bpgenc.h"

#include <stdlib.h>

static void get_plane_res(Image *img, int *pw, int *ph, int i)
{
    if (img->format == BPG_FORMAT_420 && (i == 1 || i == 2)) {
        *pw = (img->w + 1) / 2;
        *ph = (img->h + 1) / 2;
    } else if (img->format == BPG_FORMAT_422 && (i == 1 || i == 2)) {
        *pw = (img->w + 1) / 2;
        *ph = img->h;
    } else {
        *pw = img->w;
        *ph = img->h;
    }
}

void save_yuv1(Image *img, FILE *f)
{
    int c_w, c_h, i, c_count, y;

    if (img->format == BPG_FORMAT_GRAY)
        c_count = 1;
    else
        c_count = 3;
    for(i = 0; i < c_count; i++) {
        get_plane_res(img, &c_w, &c_h, i);
        for(y = 0; y < c_h; y++) {
            fwrite(img->data[i] + y * img->linesize[i], 
                   1, c_w << img->pixel_shift, f);
        }
    }
}

void save_yuv(Image *img, const char *filename)
{
    FILE *f;

    f = fopen(filename, "wb");
    if (!f) {
        fprintf(stderr, "Could not open %s\n", filename);
        exit(1);
    }
    save_yuv1(img, f);
    fclose(f);
}

HEVCEncoder* bpg_jctvc_encoder()
{
    return &jctvc_encoder;
}

