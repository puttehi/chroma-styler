package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Objective_C() Sample {
	return Sample{
		Lexer: lexers.Get(".m"),
		Code: `#import <Foundation/Foundation.h>

int main(int argc, const char * argv[]) {
  NSAutoreleasePool *pool =[[NSAutoreleasePool alloc] init];
  NSString *usage = @"Usage: please provide a string";
  if (argc < 2) {
    printf("%s\n", [usage UTF8String]);
  }
  else {
    NSString* textFromArg = [NSString stringWithUTF8String:argv[1]];
    NSString* normalizedText = [textFromArg stringByTrimmingCharactersInSet:[NSCharacterSet newlineCharacterSet]];

    if([normalizedText length] < 1){
      printf("%s\n", [usage UTF8String]);
    }
    else {
      NSString *firstChar = [[normalizedText substringToIndex:1] uppercaseString];
      NSString *remainingText = [normalizedText substringFromIndex:1];
      NSString *capitalizedText = [firstChar stringByAppendingString:remainingText];
      printf("%s\n", [capitalizedText UTF8String]);
    }
  }

  [pool drain];
  return 0;
}
`,
	}
}