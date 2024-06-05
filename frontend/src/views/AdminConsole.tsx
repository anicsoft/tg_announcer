import { ActionIcon, BoxComponentProps, Button, CloseButton, Divider, Flex, Group, Input, MantineStyleProp, Select, Stack, Switch, Text, TextInput, rem } from '@mantine/core'
import { IconAt, IconClock } from '@tabler/icons-react'
import React, { useEffect, useRef, useState } from 'react'
import RichTexInput from '../ui/RichTexInput';
import { useForm } from '@mantine/form';
import { useQuery } from '@tanstack/react-query';
import { DatePickerInput, DateTimePicker, TimeInput } from '@mantine/dates';

export default function AdminConsole() {
  const [hasPromocode, setHasPromocode] = useState(false)
  const [singleDayOffer, setSingleDayOffer] = useState(true)
  const [dateRange, setDateRange] = useState()
  const [date, setDate] = useState()
  const [startTime, setStartTime] = useState()
  const [endTime, setEndTime] = useState()
  const startTimeRef = useRef<HTMLInputElement>(null);
  const endTimeRef = useRef<HTMLInputElement>(null);


  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      title: '',
      hasPromocode: hasPromocode,
      promocode: '',
      content: '',
      dateRange: [],
      date: null,
      startTime: '',
      endTime: '',
      offerCategory: '',
      singleDayOffer: true, // DK if it is needed in the form
    },

    validate: {
      // email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
    },
  });


  const handlePromocodeSwitchChange = (isChecked) => {
    form.setValues({
      hasPromocode: isChecked
    })
    setHasPromocode(isChecked)
  }

  const handleSingleDayOfferSwitchChange = (isChecked) => {
    form.setValues({
      singleDayOffer: isChecked
    })
    setSingleDayOffer(isChecked)
  }

  const handleDateRangeChange = (value) => {
    console.log(value);

    form.setValues({
      dateRange: value
    })
    setDateRange(value)
  }

  const handleDateChange = (value) => {
    console.log(value);

    form.setValues({
      date: value
    })
    setDate(value)
  }

  const handleTimeChange = (value: Node, start: boolean) => {
    console.log(value);

    // if (start) {
    //   setStartTime(value)
    // } else {
    //   setEndTime(value)

    // }
  }

  // useEffect(() => {
  //   console.log({ ...form.getValues() }.hasPromocode);
  //   console.log({ ...form.getInputProps('hasPromocode') });

  //   setHasPromocode({ ...form.getValues() }.hasPromocode)

  //   // return () => {
  //   //   second
  //   // }
  // }, [hasPromocode])

  const { data: offerCategories } = useQuery({
    queryKey: ['offerCategories'],
    queryFn: () => fetch('http://localhost:8888/categories/offer').then((res) =>
      res.json(),
    )
  })

  const startTimePickerControl = (
    <ActionIcon variant="subtle" color="gray" onClick={() => startTimeRef.current?.showPicker()}>
      <IconClock style={{ width: rem(16), height: rem(16) }} stroke={1.5} />
    </ActionIcon>
  );
  const endTimePickerControl = (
    <ActionIcon variant="subtle" color="gray" onClick={() => endTimeRef.current?.showPicker()}>
      <IconClock style={{ width: rem(16), height: rem(16) }} stroke={1.5} />
    </ActionIcon>
  );


  const DependentFields1 = React.memo(() => {
    return (
      <>
        {hasPromocode && <TextInput
          // required
          // mt="xl"
          style={inputLabelStyles}
          label="Promocode"
          // disabled={!{ ...form.getValues() }.hasPromocode}
          key={form.key('promocode')}
          {...form.getInputProps('promocode')}
        // placeholder="Custom layout"
        // description="Write offer content here"
        // error="both below the input"
        // inputWrapperOrder={['label', 'input', 'description', 'error']}
        />}
      </>
    );
  });

  const DependentFields2 = React.memo(() => {
    return (
      <>
        {/* {!singleDayOffer && <Group
          flex={1}
        >

          <TimeInput
            label="Start time"
            withSeconds={false}
            flex={1}
            // description="Input description"
            placeholder="Start time"
            ref={startTimeRef}
            rightSection={startTimePickerControl}
            value={startTime}
            onChange={(value) => handleTimeChange(value, true)}
          />
          <TimeInput
            label="End time"
            withSeconds={false}
            flex={1}
            // description="Input description"
            placeholder="End time"
            ref={endTimeRef}
            rightSection={endTimePickerControl}
            value={endTime}
            onChange={(value) => handleTimeChange(value, false)}
          />
        </Group>} */}
        <Input.Wrapper label="Offer content">
          {/* <Input placeholder="Input inside Input.Wrapper" /> */}
          <RichTexInput content={{ ...form.getInputProps('content') }.defaultValue} onChange={handleOfferContentChange}></RichTexInput>
        </Input.Wrapper >
      </>
    );
  });

  const handleOfferContentChange = (value) => {
    form.setValues({
      content: value
    })
  }


  const inputLabelStyles: MantineStyleProp = { textAlign: "left" }
  return (
    <>
      <Text size={"lg"}>Add offer</Text>
      <form onSubmit={form.onSubmit((values) => { console.log(values), console.log(hasPromocode) })}>

        <Flex
          align={"flex-start"}
          direction={"column"}
          px="lg"
          gap={48}
        >
          <Stack
            bg="var(--mantine-color-body)"
            align="stretch"
            justify="center"
            gap="md"
          >
            <TextInput
              required
              label="Title"
              style={inputLabelStyles}
              key={form.key('title')}
              {...form.getInputProps('title')}
            // placeholder="Title"
            // description="Description below the input"
            // inputWrapperOrder={['label', 'error', 'input', 'description']}
            />
            <Select
              label="Offer type"
              placeholder="Pick value"
              required
              key={form.key('offerCategory')}
              data={offerCategories ? offerCategories.map(cat => cat.name) : []}
              {...form.getInputProps('offerCategory')}
            />
            <Switch
              // defaultChecked
              style={inputLabelStyles}
              label="Offer has a promocode"
              key={form.key('hasPromocode')}


              type='checkbox'
              checked={hasPromocode}
              // value={hasPromocode}
              onChange={(event) => handlePromocodeSwitchChange(event.currentTarget.checked)}
            // {...form.getInputProps(`hasPromocode`, { type: 'checkbox' })}
            />
            <DependentFields1></DependentFields1>

          </Stack>
          <Stack
            bg="var(--mantine-color-body)"
            align="stretch"
            justify="center"
            gap="md"
          >
            <Input.Wrapper
              required

              style={inputLabelStyles}
              // label="Some label"
              // description="Some description"
              label="Offer content"
              // placeholder="Custom layout"
              description="Write offer content here"

              key={form.key('content')}
              {...form.getInputProps('content')}
            // error="Some error"
            >
              {/* <RichTexInput></RichTexInput> */}
              <DependentFields2></DependentFields2>
              {/* <Input<any> component={TextEditor} /> */}
            </Input.Wrapper>
          </Stack>
          <Divider my="md" color='#000 ' h={1} w={"100%"} />
          <Stack
            bg="var(--mantine-color-body)"
            align="stretch"
            justify="center"
            gap="md"
            w={"100%"}
          >
            <Text size={"lg"}>Offer duration</Text>
            <DatePickerInput
              label="Pick date"
              placeholder="Pick single date"
              value={date}
              onChange={(value) => handleDateChange(value)}
            // key={form.key('date')}
            // {...form.getInputProps('date')}
            />
            {/* <DatePickerInput
              type="multiple"
              label="Pick dates"
              placeholder="Pick dates"
              value={dateRange}
              onChange={(value) => handleDateRangeChange(value)}
            /> */}
            <Switch
              // defaultChecked
              style={inputLabelStyles}
              label="Whole day"
              key={form.key('singleDayOffer')}
              defaultChecked={singleDayOffer}
              type='checkbox'
              checked={singleDayOffer}
              // value={hasPromocode}
              onChange={(event) => handleSingleDayOfferSwitchChange(event.currentTarget.checked)}
            // {...form.getInputProps(`hasPromocode`, { type: 'checkbox' })}
            />
            {!singleDayOffer && <Group
              flex={1}
            >

              <TimeInput
                label="Start time"
                withSeconds={false}
                flex={1}
                // description="Input description"
                placeholder="Start time"
                ref={startTimeRef}
                rightSection={startTimePickerControl}
                value={startTime}
                onChange={(value) => handleTimeChange(value, true)}
              />
              <TimeInput
                label="End time"
                withSeconds={false}
                flex={1}
                // description="Input description"
                placeholder="End time"
                ref={endTimeRef}
                rightSection={endTimePickerControl}
                value={endTime}
                onChange={(value) => handleTimeChange(value, false)}
              />
            </Group>}
            {/* <DependentFields2></DependentFields2> */}

          </Stack>
          <Button type="submit">Submit</Button>
        </Flex>
      </form>

    </>
  )
}
